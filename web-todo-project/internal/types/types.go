package myTypes 


// my task struct
type Task struct {
	Tid int `json:"id"`
	TDescription string `json:"description"`
}

// my task slice as local not DB 
var TaskList = make([]Task,0,5)