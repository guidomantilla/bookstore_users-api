package model

const (
	STATUS_ACTIVE     = "active"
	STATUS_NOT_ACTIVE = "inactive"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Date      string `json:"date"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}
