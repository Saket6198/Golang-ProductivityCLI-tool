package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title     string
	Completed bool
	CreatedAt time.Time
	CompletedAt time.Time
}

type Todos []Todo // new array to type Todo


func (todos *Todos) add(Title string){	// a func with a method receiver of type Todos (pointer)
	todo := Todo{	// storing a new Todo constructor in todo
		Title: Title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{}, // time.Time{} is the zero value for time.Time, nil cannot be used
	}
	*todos = append(*todos, todo)	// dereferencing so that we can append to the original slice
	fmt.Println("Todo list Added")

}

func (todos *Todos) validateIdx(index int) error {
	if index < 0 || index >= len(*todos){
		return errors.New("invalid index")
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIdx(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	fmt.Println("Todo list deleted")
	return nil
}


func (todos *Todos) toggle(index int) error{
	
	if err := todos.validateIdx(index); err != nil {
		return err
	}
	completed := (*todos)[index].Completed
	if !completed{
		(*todos)[index].Completed = true
		(*todos)[index].CompletedAt = time.Now()
		return nil
	}
	(*todos)[index].Completed = false
	(*todos)[index].CompletedAt = time.Time{}
	fmt.Println("Todo list toggled")
	return nil
}

func (todos *Todos) edit(index int, title string) error{
	if err := todos.validateIdx(index); err != nil {
		return err
	}
	(*todos)[index].Title = title
	fmt.Println("Todo list Edited")
	return nil
}

func (todos *Todos) print(){
	table := table.New(os.Stdout)	// creating a table and we use os.stdout because we want to print to the terminal
	table.SetHeaders("#", "Task", "Completed", "Created At", "Completed At")
	for index, t := range *todos{	// index keeps track of the index of the slice and t keeps current obj
		completed := "❌"
		completedAt := ""
		if t.Completed {	// working on a copy of slice
			completed = "✅"
			completedAt = t.CompletedAt.Format(time.RFC1123)
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}