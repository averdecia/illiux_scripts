package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// AddSubscriptionCommand struct is created for implementing ICommand interface
type AddSubscriptionCommand struct {
	args Args
}

// ExecuteAction implements ICommand interface
func (c *AddSubscriptionCommand) ExecuteAction(element []string) (string, error) {
	fmt.Printf("Element: %v \n", element)
	user := NewRow(element)

	resp, err := c.handleClient(user)
	if err != nil {
		return resp, err
	}

	resp2, err2 := c.addSubscription(AddSubscriptionBody, user)
	if err2 != nil {
		return resp2, err2
	}

	return resp2, nil
}

func (c AddSubscriptionCommand) handleClient(user *Row) (string, error) {

	_, err := c.getClient(user.UserID)
	if err != nil {
		fmt.Printf("User not found, Inserting: %v \n", err)
		_, err := c.insertOrUpdate(AdicionarClienteBody, user)
		if err != nil {
			return "Fail", err
		}
	} else {
		fmt.Printf("User found, Updating: %v \n", user.UserID)
		_, err := c.insertOrUpdate(ActualizarClienteBody, user)
		if err != nil {
			return "Fail", err
		}
	}
	return "Success", nil
}

func (c AddSubscriptionCommand) getClient(userid string) (string, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Minute),
	}
	rbody := fmt.Sprintf(ObtenerClienteBody,
		userid,
	)
	fmt.Println(rbody)
	request, _ := http.NewRequest("POST", c.args.Endpoint+"/dla/soap/", bytes.NewBuffer([]byte(rbody)))
	request.Header.Set("Authorization", "Basic "+c.args.AuthToken)
	request.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(request)
	return CheckResponseCode(resp, err)
}

func (c AddSubscriptionCommand) insertOrUpdate(bodyvar string, user *Row) (string, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Minute),
	}
	rbody := fmt.Sprintf(bodyvar,
		user.UserID,
		user.NombreUser,
		user.ApellidoPaterno,
		user.Email,
		user.TelefonoTelmex,
		"2", //user.FormaDePago,
	)
	fmt.Println(rbody)
	request, _ := http.NewRequest("POST", c.args.Endpoint+"/dla/soap/", bytes.NewBuffer([]byte(rbody)))
	request.Header.Set("Authorization", "Basic "+c.args.AuthToken)
	request.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(request)
	return CheckResponseCode(resp, err)
}

func (c AddSubscriptionCommand) addSubscription(bodyvar string, user *Row) (string, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Minute),
	}
	rbody := fmt.Sprintf(bodyvar,
		user.IDSubscripcion,
		user.Nombre,
		user.Item,
		user.Leyenda,
		user.UserID,
		user.FechaInicio,
		user.FechaInicioCiclo,
		user.IDPago,
		user.CodigoPromo,
	)
	fmt.Println(rbody)
	request, _ := http.NewRequest("POST", c.args.Endpoint+"/dla/soap/", bytes.NewBuffer([]byte(rbody)))
	request.Header.Set("Authorization", "Basic "+c.args.AuthToken)
	request.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(request)
	return CheckResponseCode(resp, err)
}
