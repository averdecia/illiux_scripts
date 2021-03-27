package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// DeleteCommand struct is created for implementing ICommand interface
type DeleteCommand struct {
	args Args
}

// ExecuteAction implements ICommand interface
func (c *DeleteCommand) ExecuteAction(element []string) (string, error) {
	fmt.Printf("Element: %v \n", element)
	user := NewRow(element)

	_, err := c.removeSubscription(user)

	resp2, err2 := c.removeClient(user)

	if err != nil || err2 != nil {
		return resp2, fmt.Errorf("Subs: %v, Client: %v ", err, err2)
	}

	return resp2, err
}

func (c DeleteCommand) removeSubscription(user *Row) (string, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Minute),
	}
	rbody := fmt.Sprintf(CancelarSubscription,
		user.UserID,
		user.IDSubscripcion,
	)
	request, _ := http.NewRequest("POST", c.args.Endpoint+"/dla/soap/", bytes.NewBuffer([]byte(rbody)))
	request.Header.Set("Authorization", "Basic "+c.args.AuthToken)
	request.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(request)
	return CheckResponseCode(resp, err)
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
	return CheckResponseCode(resp, err)
}
