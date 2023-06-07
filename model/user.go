package model

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	StudentId string `json:"studentId"`
	Major     string `json:"major"`
	Age       int    `json:"age"`
}
