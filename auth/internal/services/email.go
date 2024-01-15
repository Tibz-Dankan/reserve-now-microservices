package services

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Email struct {
	Recipient string `json:"recipient"`
	UserName  string `json:"name,omitempty"`
	BASE_URL  string `json:"BASE_URL,omitempty"`
}

func (e Email) sendMailRequest(URL_PATH string, body io.Reader) (*http.Response, error) {

	e.BASE_URL = "http://localhost:5000/api/v1/email"
	URL := e.BASE_URL + URL_PATH

	req, err := http.NewRequest(http.MethodPost, URL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (e Email) SendPasswordReset(resetURL string) error {

	rBody := fmt.Sprintf(`{"resetURL": "%s", "recipient": "%s", "name": "%s"}`, resetURL, e.Recipient, e.UserName)
	body := strings.NewReader(rBody)

	res, err := e.sendMailRequest("/password-reset", body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(resBody))
	}

	fmt.Printf("mail-service: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Printf("mail-service: response body: %s\n", resBody)

	return nil
}
