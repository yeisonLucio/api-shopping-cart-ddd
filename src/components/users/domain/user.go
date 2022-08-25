package domain

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
