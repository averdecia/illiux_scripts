package main

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

// DeleteCommand struct is created for implementing ICommand interface
type DeleteCommand struct {
	args Args
}

// ExecuteAction implements ICommand interface
func (c *DeleteCommand) ExecuteAction(element []string) (string, error) {
	fmt.Printf("Element: %v", element)
	user := NewRow(element)

	_, err := c.removeSubscription(user)
	if err != nil {
		fmt.Printf("Unable to remove subscription: %v", err)
	}

	resp, err2 := c.removeClient(user)
	if err2 != nil {
		return resp, err2
	}

	return resp, nil
}

func (c DeleteCommand) removeSubscription(user *Row) (string, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Minute),
	}
	rbody := fmt.Sprintf(CancelarSubscription,
		user.UserID,
		user.IDSubscripcion,
		getCurrentFormatedTime(),
	)
	request, _ := http.NewRequest("POST", c.args.Endpoint+"/dla/soap/", bytes.NewBuffer([]byte(rbody)))
	request.Header.Set("Authorization", "Basic "+c.args.AuthToken)
	request.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("Endpoint: %s, Body: %s", c.args.Endpoint, rbody)
		fmt.Printf("Cannot connect to server: %v", err)
		return "Error", err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body := buf.String()

	re := regexp.MustCompile("codigo&gt;(.*?)&lt;/codigo")
	matches := re.FindStringSubmatch(body)

	if len(matches) < 2 || matches[1] != "0" {
		return "Error", fmt.Errorf("Error code: %v", matches)
	}

	return resp.Status, nil
}

func (c DeleteCommand) removeClient(user *Row) (string, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Minute),
	}
	rbody := fmt.Sprintf(BorrarClienteBody,
		user.UserID,
	)
	request, _ := http.NewRequest("POST", c.args.Endpoint+"/dla/soap/", bytes.NewBuffer([]byte(rbody)))
	request.Header.Set("Authorization", "Basic "+c.args.AuthToken)
	request.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("Endpoint: %s, Body: %s", c.args.Endpoint, rbody)
		fmt.Printf("Cannot connect to server: %v", err)
		return "Error", err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body := buf.String()

	re := regexp.MustCompile("codigo&gt;(.*?)&lt;/codigo")
	matches := re.FindStringSubmatch(body)

	if len(matches) < 2 || matches[1] != "0" {
		return "Error", fmt.Errorf("Error code: %v", matches)
	}

	return resp.Status, nil
}

func getCurrentFormatedTime() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
}
