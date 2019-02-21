package atlas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "http://localhost:3030/v1/sms/send_sms"

// Client is a representation of a person
type Client struct {
	Username string
	Password string
}

// NewBasicAuthClient is a representation of a person
func NewBasicAuthClient(username, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

// Sms is the sms structure
type Sms struct {
	To      string `json:"to"`
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

// SendSMS is to send the SMS
func (s *Client) SendSMS(sms *Sms) ([]byte, error) {
	url := baseURL
	fmt.Println(url)
	j, err := json.Marshal(sms)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(j))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	resp, err := s.doRequest(req)
	return resp, err
}

// doRequest actually does the sendin
func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.Password)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
