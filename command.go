package main

// import (
// 	"flag"
// 	"strconv"
// 	"strings"
// )

// type CmdFlags struct {
// 	Add      string
// 	Edit     string
// 	Delete   int
// 	Complete int
// 	Display  bool
// }

// func NewCommandFlags() *CmdFlags {
// 	cf := CmdFlags{}
// 	flag.StringVar(&cf.Add, "add", "", "Add a new todo item specify title and priority (1(low)-5(high))")
// 	flag.StringVar(&cf.Edit, "edit", "", "Edit an existing todo by index &  specify title")
// 	flag.IntVar(&cf.Delete, "delete", -1, "Delete a todo item specify index")
// 	flag.IntVar(&cf.Complete, "complete", -1, "Complete a todo item specify index")
// 	flag.BoolVar(&cf.Display, "display", false, "Display all todo items")
// 	flag.Parse()
// 	return &cf
// }
// func (cf *CmdFlags) Execute(todos *Todos) {
// 	switch {
// 	case cf.Display:
// 		todos.DisplayTodos()
// 	case cf.Add != "":
// 		parts := strings.Split(cf.Add, ":")
// 		if len(parts) != 2 {
// 			println("Invalid format for add command. Use: -add title:priority")
// 			return
// 		}
// 		title := parts[0]
// 		priority, err := strconv.Atoi(parts[1])
// 		if err != nil || priority < 1 || priority > 5 {
// 			println("Invalid priority for add command. Please specify a priority between 1 and 5.")
// 			return
// 		}
// 		todos.AddTodo(title, priority, "")
// 	case cf.Edit != "":
// 		parts := strings.Split(cf.Edit, ":")
// 		if len(parts) != 2 {
// 			println("Invalid format for edit command. Use: -edit index:title")
// 			return
// 		}
// 		index, err := strconv.Atoi(parts[0])
// 		if err != nil {
// 			println("Invalid index for edit command:", err.Error())
// 			return
// 		}
// 		todos.EditTodo(index, parts[1])
// 	case cf.Delete != -1:
// 		todos.DeleteTodo(cf.Delete)
// 	case cf.Complete != -1:
// 		todos.CompleteTodo(cf.Complete)
// 	default:
// 		println("No valid command provided. Use -h for help.")
// 	}

// }
