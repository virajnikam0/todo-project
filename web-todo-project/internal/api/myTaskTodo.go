package myTask

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/virajnikam0/todo-project/web-todo-project/internal/repeatcode"
	"github.com/virajnikam0/todo-project/web-todo-project/internal/response"
	"github.com/virajnikam0/todo-project/web-todo-project/internal/types"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && r.URL.Path == "/add" {
		var myTask myTypes.Task

		// getting data from request body 
		// decode the data 
		if err := json.NewDecoder(r.Body).Decode(&myTask);err != nil {
			log.Fatalf("error in decoding : %v\n",err)
		}
		// verify data 
		// get the data from mytask like ID and description
		zTempId := myTask.Tid
		if zTempDescription := myTask.TDescription; zTempDescription == "" {
			log.Fatalf("description is empty: %v\n",zTempDescription)
		}
		zTempId = len(myTypes.TaskList)+1
		fmt.Println("Task list count: ",zTempId)
		myTask.Tid = zTempId
		myTypes.TaskList = append(myTypes.TaskList, myTask)

		 // sending the response 
		 w.Header().Set("Content-Type", "application/json") // type need to set for response

	}else{
		// setting for unvalid request
		http.Error(w,"not allowed page",http.StatusBadRequest)
	}
	defer repeat.ReadTask()
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		// take the id from url 
		tid,_ := strconv.Atoi(r.PathValue("id"))
		// check if task id present or not 
		if repeat.TaskIdPresentOrNot(tid){
			repeat.DeleteTask(tid)
			w.Header().Set("Content-Type", "application/json") 
			// response.CorrectResponse}
		}else{
			http.Error(w,"task was not present",http.StatusBadRequest)
		}
	}else{
		http.Error(w,"not allowed page",http.StatusBadRequest)
	}
	defer repeat.ReadTask()
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// check for request and method 
	if r.Method == http.MethodPut {
		// take data from URL
		tid,_ := strconv.Atoi(r.PathValue("id"))
		if repeat.TaskIdPresentOrNot(tid) {
			// oldTask := myTypes.TaskList[tid]
			newTask := myTypes.Task{}
			if err := json.NewDecoder(r.Body).Decode(&newTask);err != nil {
				http.Error(w,"error while new task",http.StatusBadRequest)
			}else{
				myTypes.TaskList[tid].TDescription = newTask.TDescription
				w.Header().Set("Content-Type", "application/json") 
				w.WriteHeader(http.StatusOK)
				res := response.CorrectResponse{
					StatusCode: http.StatusOK,
					Data: "data has been updated correctly",
				}
				if err := json.NewEncoder(w).Encode(res);err!=nil{
					log.Fatalf("error in encoding the response: %v\n",err)
				}
			}


		}else{
			http.Error(w,"task was not present",http.StatusBadRequest)
		}
	}else{
		http.Error(w,"not allowed in page",http.StatusBadRequest)
	}
	defer repeat.ReadTask()

}

func ReadTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
		res := response.CorrectResponse{
			StatusCode: http.StatusOK,
			Data: myTypes.TaskList,
		}
		if err := json.NewEncoder(w).Encode(&res);err != nil{
			http.Error(w,"Data are not able to read it",http.StatusBadRequest)
		}
	}
}