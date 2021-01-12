package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var elements = make(chan Row)
var args = getArgs()
var start = time.Now()
var filesCount = 0
var successUsersCount = 0
var successFilesLastCheckBySeconds = 0
var failedUsersCount = 0
var deleteRequestBody = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:soap="http://dla.d.iliux.com/dla/soap/">
<soapenv:Header/>
<soapenv:Body>
   <soap:borrarCliente soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
	  <id xsi:type="xsd:int">%s</id>
   </soap:borrarCliente>
</soapenv:Body>
</soapenv:Envelope>`

var addRequestBody = `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" xmlns:ns1="https://plazavip.clarodrive.com/dla/soap/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
   <SOAP-ENV:Body>
      <ns1:altaSuscripcionCliente>
         <info xsi:type="xsd:string">
            <peticion>
               <idSuscripcion><![CDATA[%s]]></idSuscripcion>
               <nombre><![CDATA[%s]]></nombre>
               <item><![CDATA[%s]]></item>
               <leyenda><![CDATA[%s]]></leyenda>
               <idCliente><![CDATA[%s]]></idCliente>
               <fechaInicio><![CDATA[%s]]></fechaInicio>
               <fechaInicioCiclo><![CDATA[%s]]></fechaInicioCiclo>
               <idPago><![CDATA[%s]]></idPago>
               <billingPeriodo><![CDATA[MES]]></billingPeriodo>
               <billingFrecuencia><![CDATA[1]]></billingFrecuencia>
               <billingCiclo><![CDATA[0]]></billingCiclo>
               <billingPrecio><![CDATA[0]]></billingPrecio>
               <codigoPromo><![CDATA[%s]]></codigoPromo>
               <trialPeriodo><![CDATA[MES]]></trialPeriodo>
               <trialFrecuencia><![CDATA[0]]></trialFrecuencia>
               <trialCiclo><![CDATA[0]]></trialCiclo>
               <trialPrecio><![CDATA[0]]></trialPrecio>
            </peticion>
         </info>
      </ns1:altaSuscripcionCliente>
   </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

// Row is used to save the csv format
type Row struct {
	UserID            string
	IDSubscripcion    string
	Nombre            string
	Item              string
	Leyenda           string
	FechaInicio       string
	FechaInicioCiclo  string
	IDPago            string
	BillingPeriodo    string
	BillingFrecuencia string
	BillingCiclo      string
	BillingPrecio     string
	CodigoPromo       string
	TrialPeriodo      string
	TrialFrecuencia   string
	TrialCiclo        string
	TrialPrecio       string
}

func (r *Row) toArray(statusCode string, message string) []string {
	return []string{
		r.UserID,
		r.IDSubscripcion,
		r.Nombre,
		r.Item,
		r.Leyenda,
		r.FechaInicio,
		r.FechaInicioCiclo,
		r.IDPago,
		r.CodigoPromo,
		statusCode,
		message,
	}
}

// Args is used to save the args been used
type Args struct {
	Command    string // First argument with file path
	Endpoint   string // Second argument with the S3 domain
	GoRoutines int    // Third argument with amount of routines
	OutputPath string // Sixth argument with the path to save the filed rows
	AuthToken  string // Basic token for wsdl
	NCDomain   string // Basic token for wsdl
	NCUser     string // Basic token for wsdl
	NCToken    string // Basic token for wsdl
}

// InstanceData is using for mapping nextcloud url data
type InstanceData struct {
	domain string
	user   string
	file   string
	auth   string
}

func (i InstanceData) buildURL() string {
	return "https://" + i.domain + "/remote.php/dav/files/" + i.user + "/Illiux/BI/" + start.Year() + "-" + start.Day() + "/" + i.file
}

var fileMapper = map[string]string{
	"add":    "estanen_cd_no_illius.csv",
	"delete": "estanen_illius_no_cd.csv",
}

func main() {
	path := downloadFile(InstanceData{
		domain: args.NCDomain,
		user:   args.NCUser,
		file:   fileMapper[args.Command],
		auth:   "Basic " + args.NCToken,
	})

	outputPointer, outputFile := getOutputWriter()

	createRoutines(args.GoRoutines, outputPointer)
	readFile(path, ",")

	waitToFinish(outputPointer, outputFile, path)
}

func getOutputWriter() (*csv.Writer, *os.File) {
	csvfile, err := os.OpenFile(args.OutputPath, os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		fmt.Printf("Failed creating file: %s \n", err)
	}
	pointer := csv.NewWriter(csvfile)

	return pointer, csvfile
}

func createRoutines(count int, outputPointer *csv.Writer) {
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		go func() {
			for {
				select {
				case element := <-elements:
					executeAction(element, outputPointer)
					filesCount++
					printProgress()
				}
			}
		}()
	}
}

func readFile(filepath string, separator string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Reading file")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), separator)
		if len(split) < 1 {
			fmt.Printf("Ignored small Element: %v \n", scanner.Text())
			continue
		}
		elements <- Row{
			UserID:           definedOrEmpty(split, 0),
			IDSubscripcion:   definedOrEmpty(split, 1),
			Nombre:           definedOrEmpty(split, 2),
			Item:             definedOrEmpty(split, 3),
			Leyenda:          definedOrEmpty(split, 4),
			FechaInicio:      definedOrEmpty(split, 5),
			FechaInicioCiclo: definedOrEmpty(split, 6),
			IDPago:           definedOrEmpty(split, 7),
			CodigoPromo:      definedOrEmpty(split, 8),
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ending reading file")
}

func getArgs() Args {
	args := os.Args[1:]

	if len(args) < 7 {
		fmt.Printf("Not enough arguments: %v\n", args)
		os.Exit(0)
	}

	if args[0] != "add" && args[0] != "delete" {
		fmt.Printf("Invalid command: %v\n", args)
		os.Exit(0)
	}

	routines, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Not int parameter routines: %v, using default 10\n", routines)
		routines = 10
	}

	return Args{
		Command:    args[0],
		Endpoint:   args[1],
		GoRoutines: routines,
		OutputPath: args[3],
		AuthToken:  args[4],
		NCDomain:   args[5],
		NCUser:     args[6],
		NCToken:    args[7],
	}
}

func printProgress() {
	seconds := time.Since(start).Seconds()
	fmt.Printf("Mean velocity: %v f/s -- Index: %v files -- Success: %v files -- Failed: %v files \n",
		math.Round(float64(filesCount)/seconds),
		filesCount,
		successUsersCount,
		failedUsersCount,
	)
}

func executeAction(element Row, outputPointer *csv.Writer) {
	resp, err := createRequest(element)
	if err != nil {
		fmt.Printf("Server error: %v", err)
		failedUsersCount++
		outputPointer.Write(element.toArray(strconv.Itoa(resp.StatusCode), err.Error()))
		outputPointer.Flush()
		return
	}

	successUsersCount++
	defer resp.Body.Close()

}

func createRequest(row Row) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Minute),
	}

	var body string
	if args.Command == "delete" {
		body = fmt.Sprintf(deleteRequestBody, row.UserID)
	} else if args.Command == "add" {
		body = fmt.Sprintf(addRequestBody,
			row.IDSubscripcion,
			row.Nombre,
			row.Item,
			row.Leyenda,
			row.UserID,
			row.FechaInicio,
			row.FechaInicioCiclo,
			row.IDPago,
			row.CodigoPromo,
		)

	}

	request, _ := http.NewRequest("POST", args.Endpoint+"/dla/soap/", bytes.NewBuffer([]byte(body)))
	request.Header.Set("Authorization", "Basic "+args.AuthToken)
	request.Header.Set("Content-Type", "text/xml")

	fmt.Printf("Endpoint: %s, Body: %s", args.Endpoint, body)

	return client.Do(request)
}

func waitToFinish(outputPointer *csv.Writer, outputFile *os.File, path string) {

	var wg sync.WaitGroup
	wg.Add(1)

	ticker := time.NewTicker(15 * time.Second)
	quit := make(chan struct{})
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				if filesCount == successFilesLastCheckBySeconds {
					printProgress()
					close(quit)
				} else {
					successFilesLastCheckBySeconds = filesCount
				}
			case <-quit:
				ticker.Stop()
				closeOutputWriter(outputPointer, outputFile, path)
				return
			}
		}
	}()
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

func definedOrEmpty(arr []string, pos int) string {
	if len(arr) >= pos+1 {
		return arr[pos]
	}
	return ""
}

func closeOutputWriter(pointer *csv.Writer, file *os.File, path string) {
	pointer.Flush()
	file.Close()
	os.Remove(path)
}

func downloadFile(instance InstanceData) string {
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
