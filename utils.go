package main

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
)

// CheckResponseCode is used to read responses from illiux
func CheckResponseCode(resp *http.Response, err error) (string, error) {
	if err != nil {
		fmt.Printf("Cannot connect to server: %v \n", err)
		return "Error", err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body := buf.String()

	fmt.Println(body)

	re := regexp.MustCompile("codigo&gt;(.*?)&lt;/codigo")
	matches := re.FindStringSubmatch(body)

	if len(matches) < 2 {
		return "Error", fmt.Errorf(body)
	}
	if matches[1] != "0" {
		return "Error", fmt.Errorf(matches[1])
	}

	return resp.Status, nil
}
