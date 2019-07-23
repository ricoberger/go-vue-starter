package mongodb

import (
	"context"
	"fmt"

	"github.com/ricoberger/go-vue-starter/pkg/model"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// CreateUser creates a new user
func (db *DB) CreateUser(u *model.User) error {
	u.ID = primitive.NewObjectID().Hex()

	_, err := db.collections.users.InsertOne(context.Background(), u)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteErrors); ok {
			if len(writeErr) == 1 && writeErr[0].Code == 11000 {
				return fmt.Errorf("email_address_already_exists")
			}
		}

		return err
	}

	return nil
}

// GetUser returns a user
func (db *DB) GetUser(id string) (*model.User, error) {
	var user model.User

	cursor := db.collections.users.FindOne(
		context.Background(),
		bson.D{primitive.E{
			Key:   "_id",
			Value: id,
		}},
	)

	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	err := cursor.Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

// GetUserByEmail returns a user by his email address
func (db *DB) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	cursor := db.collections.users.FindOne(
		context.Background(),
		bson.D{primitive.E{
			Key:   "email",
			Value: email,
		}},
	)

	if cursor.Err() != nil {

		return nil, cursor.Err()
	}

	err := cursor.Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

// SaveUser saves the given user struct
func (db *DB) SaveUser(u *model.User) error {
	cursor := db.collections.users.FindOneAndReplace(
		context.Background(),
		bson.D{primitive.E{
			Key:   "_id",
			Value: u.ID,
		}},
		u,
	)

	if cursor.Err() != nil {
		if writeErr, ok := cursor.Err().(mongo.WriteErrors); ok {
			if len(writeErr) == 1 && writeErr[0].Code == 11000 {
				return fmt.Errorf("email_address_already_exists")
			}
		}

		return cursor.Err()
	}

	return nil
}

// DeleteUser deletes the user with the given id
func (db *DB) DeleteUser(id string) error {
	cursor := db.collections.users.FindOneAndDelete(
		context.Background(),
		bson.D{primitive.E{
			Key:   "_id",
			Value: id,
		}},
	)

	if cursor.Err() != nil {
		return cursor.Err()
	}

	return nil
}
