package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// global task list
var taskList []Task

// struct of task
type Task struct {
	task   string
	isDone bool
}

// meanu screen
func OptionToChoose() int {
	userChoice := -1
	fmt.Println(`
		1. Add Task
		2. Update Task
		3. Delete Task
		4. Done and Undone task // remain 
		5. Task list
		6. exit
	`)
	fmt.Scan(&userChoice)
	return userChoice
}

func AddTask() {
	fmt.Println("enter task ")
	reader := bufio.NewReader(os.Stdin)
	task, _ := reader.ReadString('\n')

	taskList = append(taskList, Task{task: task})
}
func DeleteTask() {
	var id = -1
	ReadAllTask()
	fmt.Println("enter task id delete")
	fmt.Scan(&id)
	taskList = slices.Delete(taskList, id, id+1)
}
func UpdateTask() {
	var id = -1
	ReadAllTask()
	fmt.Println("enter task id to update")
	fmt.Scan(&id)
	oldTask := taskList[id]
	fmt.Println("Enter your modified task ")
	reader := bufio.NewReader(os.Stdin)
	newTask, _ := reader.ReadString('\n')

	fmt.Println("Your old task: ", oldTask)
	fmt.Println("Your new task: ", newTask)

	taskList[id] = Task{task: newTask}
}
func ReadAllTask() {
	for i, t := range taskList {
		fmt.Printf("ID: %v Task: %v Done: %v\n", i, t.task, t.isDone)
	}
}
func DoneAndUndoneTask() {
	ReadAllTask()
	fmt.Println("Done and undone task id ")
	var id int
	fmt.Scan(&id)
	taskList[id].isDone = !taskList[id].isDone

}

func main() {

	// func that shows options to choose
	// loop through until the user enter the exit menu
	for {
		userInput := OptionToChoose()

		if userInput != 6 {
			switch userInput {
			case 1:
				{
					AddTask()
				}
			case 2:
				{
					UpdateTask()
				}
			case 3:
				{
					DeleteTask()
				}
			case 4:
				{
					DoneAndUndoneTask()
				}
			case 5:
				{
					ReadAllTask()
				}
			}
		} else {
			break
		}
	}

}
