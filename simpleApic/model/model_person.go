package model

type Person struct {
	ID        int  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (e *Person) TableName() string {
	return "person"
}