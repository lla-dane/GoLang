package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Text      string
	Completed bool
}

func main() {
	tasks := []Task{}

	for {
		showMenu()
		option := getUserInput("Enter your choice:")

		switch option {
		case "1":
			showTasks(tasks)
		case "2":
			addTasks(&tasks)
		case "3":
			markTaskCompleted(&tasks)
		case "4":
			saveTasksToFile(tasks)
		case "5":
			fmt.Println("Exiting the TODO application")
			return

		default:
			fmt.Println("Invalid choice, please try again")
		}
	}
}

func showMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Show Tasks")
	fmt.Println("2. Add Task")
	fmt.Println("3. Mark Task as completed")
	fmt.Println("4. Save task to the file")
	fmt.Println("5. Exit")
}

func getUserInput(promt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(promt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No Tasks available")
		return
	}
	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Text)
	}
}

func addTasks(tasks *[]Task) {
	taskText := getUserInput("Enter task description: ")
	*tasks = append(*tasks, Task{Text: taskText})
	fmt.Println("Task added.")
}

func markTaskCompleted(tasks *[]Task) {
	showTasks(*tasks)
	taskIndexStr := getUserInput("Enter task number to mark as completed: ")
	taskIndex, err := strconv.Atoi(taskIndexStr)
	if err != nil || taskIndex < 1 || taskIndex > len(*tasks) {
		fmt.Println("Invalid task number. Please try again.")
		return
	}
	(*tasks)[taskIndex-1].Completed = true
	fmt.Println("Task marked as completed.")
}

func saveTasksToFile(tasks []Task) {
	file, err := os.Create("tasks.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		file.WriteString(fmt.Sprintf("[%s] %s\n", status, task.Text))
	}
	fmt.Println("Tasks saved to file 'tasks.txt'.")
}
