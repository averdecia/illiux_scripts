package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// MailsCommand struct is created for implementing ICommand interface
type MailsCommand struct {
	args Args
}

// ExecuteAction implements ICommand interface
func (c *MailsCommand) ExecuteAction(element []string) (string, error) {
	fmt.Printf("Element: %v \n", element)
	user := NewMailRow(element)
	if user.Email == "" {
		return "Empty", nil
	}
	if user.GamificationID == "" {
		user.GamificationID = user.Email
	}

	resp, err := c.sendEmail(user)
	if err != nil {
		fmt.Printf("Unable to send the email: %v \n", err)
	}

	return resp, nil
}

func (c MailsCommand) sendEmail(user *MailRow) (string, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Minute),
	}
	rbody := fmt.Sprintf(SendEmailBody,
		c.args.NCDomain,     //key
		user.Email,          //email
		user.GamificationID, //gamification
		user.Email,          //email
	)
	request, _ := http.NewRequest("POST", c.args.Endpoint+"/v1/queue/mail/send?authpn="+c.args.NCUser+"&authpt="+c.args.NCToken,
		bytes.NewBuffer([]byte(rbody)))
	// request.Header.Set("Authorization", "Basic "+c.args.AuthToken)
	request.Header.Set("Content-Type", "application/json")
	fmt.Printf("Endpoint: %s, Body: %s \n", c.args.Endpoint, rbody)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("Endpoint: %s, Body: %s \n", c.args.Endpoint, rbody)
		fmt.Printf("Cannot connect to server: %v \n", err)
		return "Error", err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	body := buf.Bytes()

	r := bytes.NewReader(body)
	decoder := json.NewDecoder(r)

	response := &EngResponse{}
	err = decoder.Decode(response)

	if err != nil {
		return "Error", err
	}

	return response.Status, nil
}
