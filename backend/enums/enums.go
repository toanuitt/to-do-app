package enums

type TaskPriority int

const (
	Low TaskPriority = iota +1
	Medium
	High
)

type TaskCategory int 

const (
	Work TaskCategory = iota + 1
	Study
	Home
)