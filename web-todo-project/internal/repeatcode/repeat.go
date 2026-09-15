package repeat

import (
	"fmt"
	"slices"

	myTypes "github.com/virajnikam0/todo-project/web-todo-project/internal/types"
)

func TaskIdPresentOrNot(tid int) bool {
	// loop to list
	for i,_ := range myTypes.TaskList{
		if myTypes.TaskList[i].Tid == tid {
			return true
		}
	}
	return false
}

// delete the task from list
func DeleteTask(tid int)(bool){
	oldLen := len(myTypes.TaskList)
	myTypes.TaskList = slices.Delete(myTypes.TaskList,tid,tid+1)
	newLen := len(myTypes.TaskList)

	if newLen - oldLen == 1 {
		return true
	}else{
		return false
	}
} 


// read the task for only backend 
func ReadTask(){
	for i,d := range myTypes.TaskList{
		println("Task id: ",d.Tid," Task description: ",d.TDescription," index: ",i)
	}
	fmt.Println("====================================")
}