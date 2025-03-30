package models 

import (
	"backend/enums"
	"time"
)

type Task struct {
	ID 			int 				`json:"id"`
	Title 		string				`json:"title"`
	Description string				`json:"description"`
	Priority 	enums.TaskPriority  `json:"priority"`
	Category 	enums.TaskCategory  `json:"category"`
	DueDate 	time.Time			`json:"dueDate"`
	Completed 	bool				`json:"completed"`
}
