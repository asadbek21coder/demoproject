package entities

type Book struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Price     int    `json:"price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


