package entities

type Book struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Author    int    `json:"author"`
	Price     int    `json:"price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetBookByIdRes struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         *Book  `json:"body"`
}

type GetAllBooks struct {
	ErrorCode    int     `json:"error_code"`
	ErrorMessage string  `json:"error_message"`
	Body         []*Book `json:"body"`
}

type BookRes struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         *Book  `json:"body"`
}

type CreateBookReq struct {
	Name   string `json:"name"`
	Author int    `json:"author"`
	Price  int    `json:"price"`
}

type UpdateBookReq struct {
	Name   string `json:"name"`
	Author int    `json:"author"`
	Price  int    `json:"price"`
}

type DeleteBookRes struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
