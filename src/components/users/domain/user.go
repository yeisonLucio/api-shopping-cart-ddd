package user_domain

import "encoding/json"

type User struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"Password,omitempty"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Phone        string `json:"phone"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	UserType     string `json:"user_type"`
}

type LoginCredentials struct {
	UserName string
	Password string
}

func (p *User) Parse(structure interface{}) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &structure)
}
