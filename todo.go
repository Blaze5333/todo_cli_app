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
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
	Priority    int        `json:"priority"`
	UserName    string     `json:"username"`
}

type Data struct {
	Username string `json:"username"`
	Todos    []Todo `json:"todos"`
}
type Datas []Data

func (data *Datas) AddTodo(title string, priority int, userName string) {

	newTodo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
		Priority:    priority,
		UserName:    userName,
	}

	for index, todo := range *data {
		if todo.Username == userName {
			todo.Todos = append(todo.Todos, newTodo)
			addedData := Data{
				Username: todo.Username,
				Todos:    todo.Todos,
			}
			(*data)[index] = addedData
			return
		}
	}
	*data = append(*data, Data{
		Username: userName,
		Todos:    []Todo{newTodo},
	})
}
func (data *Datas) validateID(id int, username string) (int, []Todo, error) {
	for index, todo := range *data {
		if todo.Username == username {
			if id < 0 || id >= len(todo.Todos) {
				err := errors.New("invalid ID")
				fmt.Println(err)
				return -1, nil, err
			}
			return index, todo.Todos, nil
		}
	}
	return -1, nil, errors.New("user not found")
}
func (todo *Datas) DeleteTodo(id int, username string) error {
	index, todos, err := todo.validateID(id, username)
	if err != nil {
		return err
	}
	todos = append(todos[:id], todos[id+1:]...)
	(*todo)[index].Todos = todos
	return nil
}
func (todo *Datas) EditTodo(id int, title string, username string) error {
	index, todos, err := todo.validateID(id, username)
	if err != nil {
		return err
	}
	(todos)[id].Title = title
	(*todo)[index].Todos = todos
	return nil
}
func (todo *Datas) CompleteTodo(id int, username string) error {
	index, todos, err := todo.validateID(id, username)
	if err != nil {

		return err
	}
	(todos)[id].Completed = true
	currentTime := time.Now()
	(todos)[id].CompletedAt = &currentTime
	(*todo)[index].Todos = todos
	return nil
}
func (todo *Datas) DisplayTodos(username string) {
	table := table.New(os.Stdout)
	table.SetHeaders("ID", "Title", "Priority", "Completed", "Created At", "Completed At")
	var td []Todo
	for _, data := range *todo {
		if data.Username != username {
			continue
		}
		td = data.Todos
		break
	}
	for i, t := range td {
		completedAt := "-"
		if t.CompletedAt != nil {
			completedAt = t.CompletedAt.Format("2006-01-02, 3:04 PM")
		}
		completed := "❌"
		if t.Completed {
			completed = "✅"
		}
		var color string
		switch t.Priority {
		case 1:
			color = "🟢"
		case 2:
			color = "🔵"
		case 3:
			color = "🟡"
		case 4:
			color = "🔴"
		default:
			color = "⚪️"
		}

		table.AddRow(strconv.Itoa(i), t.Title, color, completed, t.CreatedAt.Format("2006-01-02, 3:04 PM"), completedAt)
	}
	table.Render()
}
