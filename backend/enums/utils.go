package enums

// conver taskpriority to string 
func (p TaskPriority) String() string {
	return [...]string{"Low", "Medium", "High"}[p-1]
}

// convert string to taskpriority
func ParseTaskPriority(s string) TaskPriority {
	switch s {
	case "Low":
		return Low
	case "Medium":
		return Medium
	case "High":
		return High
	default:
		return Low
	}
}

// convert taskcategory to string
func (c TaskCategory) String() string {
	return [...]string{"Work", "Study", "Home"}[c-1]
}

// convert string to taskcategory
func ParseTaskCategory(s string) TaskCategory {
	switch s {
	case "Work":
		return Work
	case "Study":
		return Study
	case "Home":
		return Home
	default:
		return Others
	}
}


