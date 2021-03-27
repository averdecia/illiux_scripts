package main

import (
	"fmt"
	"os"

	command "github.com/averdecia/script_command"
)

var fileMapper = map[string]string{
	"add":    "estanen_cd_no_illius.csv",
	"delete": "estanen_illius_no_cd.csv",
}

func main() {
	args := GetArgs(os.Args[1:])
	path := ""

	var mycommand command.ICommand
	switch args.Command {
	case string(Add):
		mycommand = &AddSubscriptionCommand{
			args: args,
		}
		path = Download(InstanceData{
			domain: args.NCDomain,
			user:   args.NCUser,
			file:   fileMapper[args.Command],
			auth:   "Basic " + args.NCToken,
		})
	case string(Delete):
		mycommand = &DeleteCommand{
			args: args,
		}
		path = Download(InstanceData{
			domain: args.NCDomain,
			user:   args.NCUser,
			file:   fileMapper[args.Command],
			auth:   "Basic " + args.NCToken,
		})
	case string(Mail):
		mycommand = &MailsCommand{
			args: args,
		}
		path = args.AuthToken // path to file
	default:
		fmt.Printf("Invalid command: %v\n", args.Command)
		os.Exit(0)
	}

	command.RunProcess(mycommand, args.GoRoutines, path, args.OutputPath, 60)

}
