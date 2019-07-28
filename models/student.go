package models

type Student struct {
	EmployeeID int    `json:"employeeID"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Status     bool   `json:"status"`
}

// func (Administrator) TableName() string {
// 	return "administrator_list"
// }
