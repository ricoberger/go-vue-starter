package mail

import (
	"bytes"
	"html/template"
	"net/smtp"

	"github.com/ricoberger/go-vue-starter/pkg/static"
)

// Config represents the mail configuration
type Config struct {
	Identity string            `yaml:"identity"`
	Username string            `yaml:"username"`
	Password string            `yaml:"password"`
	Host     string            `yaml:"host"`
	Addr     string            `yaml:"addr"`
	From     string            `yaml:"from"`
	HTML     bool              `yaml:"html"`
	Subjects map[string]string `yaml:"subjects"`
}

// Client implements a mail client
type Client struct {
	auth     smtp.Auth
	addr     string
	from     string
	html     bool
	subjects map[string]string
}

// NewClient return a new mail
func NewClient(config *Config) *Client {
	return &Client{
		auth:     smtp.PlainAuth(config.Identity, config.Username, config.Password, config.Host),
		addr:     config.Addr,
		from:     config.From,
		html:     config.HTML,
		subjects: config.Subjects,
	}
}

// Send sends an email message to the provided address and with the provided subject
func (c *Client) Send(to, email string, data interface{}) error {
	var mime string
	var templateFile string

	if c.html {
		mime = "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
		templateFile = "/web/mail/" + email + ".html"
	} else {
		mime = "MIME-version: 1.0;\r\nContent-Type: text/plain; charset=\"UTF-8\";\r\n\r\n"
		templateFile = "/web/mail/" + email + ".txt"
	}

	subject := "Subject: " + c.subjects[email] + "\r\n"
	fromHeader := "From: " + c.from + "\r\n"
	toHeader := "To: " + to + "\r\n"

	body, err := c.parseTemplate(templateFile, data)
	if err != nil {
		return err
	}

	msg := []byte(subject + fromHeader + toHeader + mime + "\r\n" + body)

	err = smtp.SendMail(
		c.addr,
		c.auth,
		c.from,
		[]string{to},
		[]byte(msg),
	)

	return err
}

// parseTemplate parses a html template file
func (c *Client) parseTemplate(templateFile string, data interface{}) (string, error) {
	file, err := static.FSString(false, templateFile)
	if err != nil {
		return "", err
	}

	t, err := template.New(templateFile).Parse(file)
	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
