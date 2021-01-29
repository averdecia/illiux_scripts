package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ValidCommands will save all type of commands
type ValidCommands string

const (
	// Add will create a new subscription
	Add ValidCommands = "add"
	// Delete will remove a client and his subs
	Delete ValidCommands = "delete"
)

func (e ValidCommands) String() string {
	commands := [...]string{"add", "delete"}
	x := string(e)
	for _, v := range commands {
		if strings.ToLower(v) == strings.ToLower(x) {
			return strings.ToLower(v)
		}
	}
	return ""
}

// Row is used to save the csv format
type Row struct {
	UserID           string
	IDSubscripcion   string
	Nombre           string
	Item             string
	Leyenda          string
	FechaInicio      string
	FechaInicioCiclo string
	IDPago           string
	CodigoPromo      string
	NombreUser       string
	ApellidoPaterno  string
	Email            string
	TelefonoTelmex   string
	FormaDePago      string
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
		r.NombreUser,
		r.ApellidoPaterno,
		r.Email,
		r.TelefonoTelmex,
		r.FormaDePago,
		statusCode,
		message,
	}
}

// NewRow create and object from a string array
func NewRow(element []string) *Row {
	return &Row{
		UserID:           element[0],
		IDSubscripcion:   element[1],
		Nombre:           element[2],
		Item:             element[3],
		Leyenda:          element[4],
		FechaInicio:      element[5],
		FechaInicioCiclo: element[6],
		IDPago:           element[7],
		CodigoPromo:      element[8],
		NombreUser:       element[9],
		ApellidoPaterno:  element[10],
		Email:            element[11],
		TelefonoTelmex:   element[12],
		FormaDePago:      element[13],
	}
}

// Args is used to save the args been used
type Args struct {
	Command    string // First argument with file path
	Endpoint   string // Second argument with the WSDL domain
	GoRoutines int    // Third argument with amount of routines
	OutputPath string // Argument with the path to save the filed rows
	AuthToken  string // Basic token for wsdl
	NCDomain   string // Domain for NC instance
	NCUser     string // user for NC instance
	NCToken    string // Basic token for wsdl
}

// GetArgs will retrieve the command line arguments
func GetArgs(args []string) Args {
	if len(args) < 7 {
		fmt.Printf("Not enough arguments: %v\n", args)
		os.Exit(0)
	}
	routines, _ := strconv.Atoi(args[2])

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
