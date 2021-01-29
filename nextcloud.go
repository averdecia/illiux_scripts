package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// InstanceData is using for mapping nextcloud url data
type InstanceData struct {
	domain string
	user   string
	file   string
	auth   string
}

func (i InstanceData) buildURL() string {
	year, month, _ := time.Now().Date()
	return "https://" + i.domain + "/remote.php/dav/files/" + i.user + "/Illiux/BI/" + strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "/" + i.file
}

// Download method will be used to retrieve a file from claro drive
func Download(instance InstanceData) string {
	client := http.Client{
		Timeout: time.Duration(5 * time.Minute),
	}

	request, err := http.NewRequest("GET", instance.buildURL(), nil)
	request.Header.Set("Authorization", instance.auth)
	request.Header.Set("OCS-APIRequest", "true")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("Cannot connect to server: %v", err)
	} else {
		defer resp.Body.Close()
	}
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		fmt.Printf("Cannot create temporary file: %v", err)
	}

	// Write the body to file
	_, err = io.Copy(tmpFile, resp.Body)

	return tmpFile.Name()
}
