package tasks

import (
	"fmt"
	"sort"

	"github.com/selamde/task-manager/types"
)

func DisplayTasks() {
	fmt.Println("-----------List Of Tasks-----------")

	tasks := []types.Task{
		 {
        Title: "Learn Go slices",
        Completed: false,
    },
    {
        Title: "Build Task Manager CLI",
        Completed: false,
    },
    {
        Title: "Study Gin framework",
        Completed: false,
    },
    {
        Title: "Connect PostgreSQL",
        Completed: false,
    },
    {
        Title: "Build JWT auth",
        Completed: false,
    },
    {
        Title: "Deploy Go API",
        Completed: false,
    },
	}

for index, task := range tasks{
	taskIndex := index +1
	taskTitle := task.Title
	isCompelted := task.Completed

	fmt.Printf("%d. Task_Name: %s -- Completed: %t \n", taskIndex, taskTitle, isCompelted)
}

//call

tasks = AssignIdentificationNumber(tasks)

fmt.Println("---------After adding the taskId-----")
for index, task := range tasks{
	taskIndex := index +1
	taskTitle := task.Title
	isCompelted := task.Completed
	taskId := fmt.Sprintf("Task - %04d ", taskIndex)
	tasks[index].ID = taskId

	fmt.Printf("%d. Task_Name: %s -- Completed: %t  Task_Id: %s \n", taskIndex, taskTitle, isCompelted, taskId)
}




}

func AssignIdentificationNumber(tasks []types.Task) []types.Task{


	sort.Slice(tasks, func(i, j int) bool{
return tasks[i].Title < tasks[j].Title
	})

	fmt.Println("--------After sorting-----")
for index, task := range tasks{
	taskIndex := index +1
	taskTitle := task.Title
	isCompelted := task.Completed
	taskId := fmt.Sprintf("Task - %04d ", taskIndex)
	tasks[index].ID = taskId

	fmt.Printf("%d. Task_Name: %s -- Completed: %t  Task_Id: %s \n", taskIndex, taskTitle, isCompelted, taskId)
}




	fmt.Println("------Adding the task ID--------")
for index, task := range tasks{
	taskIndex := index +1
	taskTitle := task.Title
	isCompelted := task.Completed
	taskId := fmt.Sprintf("Task - %04d ", taskIndex)
	tasks[index].ID = taskId

	fmt.Printf("%d. Task_Name: %s -- Completed: %t  Task_Id: %s \n", taskIndex, taskTitle, isCompelted, taskId)
}


	return tasks

}

func AddTask(tasks []types.Task, title string) []types.Task{
 newTask := types.Task{
	Title: title,
	Completed: false,
 }

 tasks = append(tasks, newTask)


 return AssignIdentificationNumber(tasks)



}