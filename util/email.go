package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type EmailService struct {
	apiKey    string
	parterKey string
	LiveEMail bool
}

type EmailServiceOpts struct {
	APIKey     string
	PartnerKey string
	LiveEMail  bool
}

func NewEmailService(opts EmailServiceOpts) (*EmailService, error) {
	return &EmailService{
		apiKey:    opts.APIKey,
		parterKey: opts.PartnerKey,
		LiveEMail: opts.LiveEMail,
	}, nil

}

type OutgoingEmailOpts struct {
	Sender     Sender `json:"sender"`
	To         []To   `json:"to"`
	TemplateID int    `json:"templateId"`
	Params     Params `json:"params"`
}
type Sender struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type To struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type Params struct {
	FirstName string `json:"first_name"`
	Lname     string `json:"last_name"`
}

func (svc EmailService) SendEmail(ctx context.Context, opts OutgoingEmailOpts) bool {

	url := "https://api.sendinblue.com/v3/smtp/email"
	method := "POST"

	byteMessage, err := json.Marshal(opts)
	if err != nil {
		fmt.Println("unable to marshall", err)
	}
	payload := strings.NewReader(string(byteMessage))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Add("api-key", svc.apiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(string(body))
	return true
}
