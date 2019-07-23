package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/ricoberger/go-vue-starter/pkg/api/response"
	"github.com/ricoberger/go-vue-starter/pkg/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// loginHandler handles user authentication
func (a *API) userLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Validate user input
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if u.Email == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Email address is missing")
		return
	} else if u.Password == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Password is missing")
		return
	}

	// Always use the lower case email address
	u.Email = strings.ToLower(u.Email)

	// Get the user database entry
	user, err := a.db.GetUserByEmail(u.Email)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if user == nil {
		response.Errorf(w, r, err, http.StatusBadRequest, "Invalid email address or password")
		return
	}

	// Check the password
	if !user.MatchPassword(u.Password) {
		response.Errorf(w, r, err, http.StatusBadRequest, "Invalid email address or password")
		return
	}

	// Check if the verify token is set
	if user.VerifyToken != "" {
		response.Errorf(w, r, err, http.StatusBadRequest, "Account is not verifyed")
		return
	}

	// Omit password, reset and verify token
	user.Password = ""
	user.ResetPasswordToken = ""
	user.VerifyToken = ""

	// Create jwt token
	user.Token, err = a.createJWT(jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.Write(w, r, user)
	return
}

// signupHandler handles user sign up
func (a *API) userSignupHandler(w http.ResponseWriter, r *http.Request) {
	// Validate user input
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if u.Name == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Name is missing")
		return
	} else if u.Email == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Email address is missing")
		return
	} else if u.Password == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Password is missing")
		return
	}

	// Always use the lower case email address
	u.Email = strings.ToLower(u.Email)

	// Hash the user password
	err = u.HashPassword()
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Create verify token
	u.VerifyToken = a.createRandString(32)

	// Save user to database
	err = a.db.CreateUser(&u)
	if err != nil {
		if err.Error() == "email_address_already_exists" {
			response.Errorf(w, r, err, http.StatusBadRequest, "Email address already exists")
			return
		}

		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Send welcome email
	templateData := struct {
		Name string
		URL  string
	}{
		Name: u.Name,
		URL:  a.config.Domain + "/verify/" + u.ID + "/" + u.VerifyToken,
	}

	go func(email string, data interface{}) {
		err = a.mail.Send(email, "welcome", data)
		if err != nil {
			logrus.WithError(err).Error("Could not send welcome email")
		}
	}(u.Email, templateData)

	// Omit password, reset and verify token
	u.Password = ""
	u.ResetPasswordToken = ""
	u.VerifyToken = ""

	response.Write(w, r, u)
	return
}

func (a *API) userUpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get jwt claims
	claims, ok := r.Context().Value(ContextJWTKey).(jwt.MapClaims)
	if !ok {
		response.Errorf(w, r, nil, http.StatusUnauthorized, "Unauthorized")
	}

	// Get user id
	id, err := a.getID(claims)
	if err != nil {
		response.Errorf(w, r, nil, http.StatusUnauthorized, "Unauthorized")
	}

	// Validate user input
	var u model.User
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if u.Email == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Email address is missing")
		return
	} else if u.Name == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Name is missing")
		return
	}

	// Get the user
	user, err := a.db.GetUser(id)
	if err != nil || user == nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Change name
	user.Name = u.Name

	if user.Email != u.Email {
		oldEmail := user.Email

		// Change email address
		user.EmailBackup = append(user.EmailBackup, user.Email)
		user.Email = u.Email
		user.VerifyToken = a.createRandString(32)

		// Send verification email
		templateData0 := struct {
			Name string
			URL  string
		}{
			Name: user.Name,
			URL:  a.config.Domain + "/verify/" + user.ID + "/" + user.VerifyToken,
		}

		err = a.mail.Send(user.Email, "email-verification", templateData0)
		if err != nil {
			response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Send notification email
		templateData1 := struct {
			Name     string
			NewEmail string
		}{
			Name:     user.Name,
			NewEmail: user.Email,
		}

		err = a.mail.Send(oldEmail, "email-changed", templateData1)
		if err != nil {
			response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		err = a.db.SaveUser(user)
		if err != nil {
			if err.Error() == "email_address_already_exists" {
				response.Errorf(w, r, err, http.StatusBadRequest, "Email address already exists")
				return
			}

			response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
			return
		}
	} else {
		// Save the user
		err = a.db.SaveUser(user)
		if err != nil {
			response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}

	// Omit password and reset token
	user.Password = ""
	user.ResetPasswordToken = ""
	user.VerifyToken = ""

	// Create jwt token
	user.Token, err = a.createJWT(jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.Write(w, r, user)
	return
}

func (a *API) userProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get jwt claims
	claims, ok := r.Context().Value(ContextJWTKey).(jwt.MapClaims)
	if !ok {
		response.Errorf(w, r, nil, http.StatusInternalServerError, "Internal Server Error")
	}

	// Get user id
	id, err := a.getID(claims)
	if err != nil {
		response.Errorf(w, r, nil, http.StatusUnauthorized, "Unauthorized")
	}

	// Get the user
	user, err := a.db.GetUser(id)
	if err != nil || user == nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Omit password and reset token
	user.Password = ""
	user.ResetPasswordToken = ""
	user.VerifyToken = ""

	response.Write(w, r, user)
	return
}

func (a *API) userVerifyHandler(w http.ResponseWriter, r *http.Request) {
	// Validate user input
	vars := mux.Vars(r)
	token := vars["token"]
	id := vars["id"]

	// Get user
	user, err := a.db.GetUser(id)
	if err != nil || user == nil {
		response.Errorf(w, r, err, http.StatusBadRequest, "Invalid account id")
		return
	}

	// Verify the account verification token
	if user.VerifyToken != token {
		response.Errorf(w, r, err, http.StatusBadRequest, "Invalid account verification token")
		return
	}

	// Delete verification token
	user.VerifyToken = ""

	// Save the user
	err = a.db.SaveUser(user)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.Write(w, r, nil)
	return
}

func (a *API) userResendVerificationMail(w http.ResponseWriter, r *http.Request) {
	// Validate user input
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if u.ID == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Account id is missing")
		return
	}

	// Get the user
	user, err := a.db.GetUser(u.ID)
	if err != nil || user == nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Send welcome email
	templateData := struct {
		Name string
		URL  string
	}{
		Name: user.Name,
		URL:  a.config.Domain + "/verify/" + user.ID + "/" + user.VerifyToken,
	}

	err = a.mail.Send(user.Email, "welcome", templateData)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Verification email could not be send")
		return
	}

	// Omit password, reset and verify token
	user.Password = ""
	user.ResetPasswordToken = ""
	user.VerifyToken = ""

	response.Write(w, r, u)
	return
}

func (a *API) forgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Validate user input
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if u.Email == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Email address is missing")
		return
	}

	// Always use the lower case email address
	u.Email = strings.ToLower(u.Email)

	// Get user
	user, err := a.db.GetUserByEmail(u.Email)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if user == nil {
		response.Errorf(w, r, err, http.StatusBadRequest, "Invalid email address")
		return
	}

	// Create reset password token
	user.ResetPasswordToken = a.createRandString(32)

	// Save the user
	err = a.db.SaveUser(user)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Send reset email
	templateData := struct {
		Name string
		URL  string
	}{
		Name: u.Name,
		URL:  a.config.Domain + "/reset-password/" + user.ID + "/" + user.ResetPasswordToken,
	}

	err = a.mail.Send(u.Email, "reset-password", templateData)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Could not send reset password email, please try it again")
		return
	}

	response.Write(w, r, nil)
	return
}

func (a *API) resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Validate user input
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if u.ID == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "ID address is missing")
		return
	} else if u.ResetPasswordToken == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Reset password token is missing")
		return
	} else if u.Password == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "New password is missing")
		return
	}

	// Get user
	user, err := a.db.GetUser(u.ID)
	if err != nil || user == nil {
		response.Errorf(w, r, err, http.StatusBadRequest, "Invalid email address")
		return
	}

	// Check reset token
	if user.ResetPasswordToken != u.ResetPasswordToken {
		response.Errorf(w, r, err, http.StatusBadRequest, "Invalid reset password token")
		return
	}

	// Set new values for reset password token and password
	user.ResetPasswordToken = ""
	user.Password = u.Password

	// Hash the new password
	err = user.HashPassword()
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Save the user
	err = a.db.SaveUser(user)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.Write(w, r, nil)
	return
}
