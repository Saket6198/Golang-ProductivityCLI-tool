package main

// import (
// 	"fmt"
// 	"time"
// )

func main(){
	// fmt.Println("Initializing TODO CLI... [SYSTEM ONLINE]")
	// time.Sleep(1 * time.Second)
	// fmt.Println("Welcome to the Coder Army TODO Manager - Your tasks, your mission.")
	// time.Sleep(1 * time.Second)
	// fmt.Println("Loading modules... [OK]")
	// time.Sleep(1 * time.Second)
	// fmt.Println("Ready to conquer your tasks. Let's hack productivity!")

	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlag := NewCmdFlags()
	cmdFlag.Execute(&todos)
	storage.Save(todos)
}