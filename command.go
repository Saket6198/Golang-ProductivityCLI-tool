package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Toggle int
	Edit   string
	List   bool
}

func NewCmdFlags() *CmdFlags {
	if len(os.Args) == 1 {
		fmt.Println("ProdCLI version 1.0.0")
		fmt.Println("Type go run . -help for more information")
		os.Exit(0)
	}
	cmd := CmdFlags{}

	flag.StringVar(&cmd.Add, "add", "", "Add a new todo")
	flag.IntVar(&cmd.Del, "del", -1, "Specify a index to Delete a task")
	flag.StringVar(&cmd.Edit, "edit", "", "Edit a task by index & specify a new title")
	flag.IntVar(&cmd.Toggle, "toggle", -1, "Specify a index to Toggle a task")
	flag.BoolVar(&cmd.List, "list", false, "List all todos")
	flag.Parse()
	return &cmd
}

func (cmd *CmdFlags) Execute(todos *Todos){
	switch {
		case cmd.List:
			todos.print()
		case cmd.Add != "":
			todos.add(cmd.Add)
		case cmd.Edit != "":
			parts := strings.SplitN(cmd.Edit, " ", 2)
			if len(parts) != 2{
				fmt.Println("Error, Invalid format for edit. Specify index and title")
				os.Exit(1)
			}	
			index, err := strconv.Atoi(parts[0])
			if err != nil{
				fmt.Println("Error, Invalid index")
				os.Exit(1)
			}
			todos.edit(index, parts[1])
		case cmd.Del != -1:
			todos.delete(cmd.Del)
		case cmd.Toggle != -1:
			todos.toggle(cmd.Toggle)
		default:
			fmt.Println("Error, Invalid command")
	}
}