package main

import "fmt"

type Todo struct {
	todo   string
	isDone bool
}

func createNewTodo(task string, isDone bool) Todo {
	return Todo{
		todo:   task,
		isDone: isDone,
	}
}

func createTodo(todoList *[]Todo, todo string, isDone bool) {
	newTodo := createNewTodo(todo, isDone)
	*todoList = append(*todoList, newTodo)
}

func markAsCompleted(todoList *[]Todo, todo string, isCompleted bool) *[]Todo {
	for i, oldTodo := range *todoList {
		if todo == oldTodo.todo {
			(*todoList)[i].isDone = isCompleted
			fmt.Printf("[%s] Marked as Completed: %t\n", todo, isCompleted)
			return todoList
		}
	}
	fmt.Printf("[%s] Selected todo not found\n", todo)
	return todoList
}

func displayTodo(todoList *[]Todo, todoTitle string) Todo {
	for i, todo := range *todoList {
		if todo.todo == todoTitle {

			fmt.Printf("%d. [%s] Completed: %t\n", i+1, todo.todo, todo.isDone)
			return todo
		}
	}
	fmt.Printf("[%s] Todo not found\n", todoTitle)
	return Todo{}
}

func mappedTodo(todoList []Todo) {
	for i, todo := range todoList {
		fmt.Printf("%d. %s (Completed: %t)\n", i+1, todo.todo, todo.isDone)
	}
}
func main() {

	var todoList []Todo
	createTodo(&todoList, "Go to the gym", false)
	createTodo(&todoList, "Go to the prayer", false)
	createTodo(&todoList, "Go to the eat", false)

	markAsCompleted(&todoList, "Go to the prayer", true)
	markAsCompleted(&todoList, "Go to the gym", true)
	markAsCompleted(&todoList, "Go to the eat", true)

	mappedTodo(todoList)

	displayTodo(&todoList, "Go to the eat")
}
