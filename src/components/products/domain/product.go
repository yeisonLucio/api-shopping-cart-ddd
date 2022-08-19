package domain

import "encoding/json"

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
}

func (p *Product) Parse(structure interface{}) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &structure)
}
