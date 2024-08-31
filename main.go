package main

import (
	"fmt"
)

type Task struct {
	TaskName string
	Completed bool
}

var tasks []Task

func addTask(task string) {
	newTask := Task{TaskName: task, Completed: false}
	tasks = append(tasks, newTask)
	fmt.Println("Task Added")
}

func listTasks() {
	for i, task := range tasks {
		var status string = "n"
		if task.Completed {
			status = "d"
		}

		fmt.Printf("%d. %s [%s]\n", i+1, task.TaskName, status)
	}
}

func markCompleted(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks[index - 1].Completed = true
		fmt.Println("Task Marked as Completed")
	} else {
		fmt.Println("Invalid Index")
	}
}

func editTask(index int, newString string) {
	if index >= 1 && index <= len(tasks) {
		// tasks[index - 1].TaskName = Task()
		fmt.Println("Task edited successfully")
	} else {
		fmt.Println("Invalid Index")
	}
}

func deleteTask(index int) {
	
}

func main() {

}