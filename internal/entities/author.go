package entities

type Author struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetAuthorByIdRes struct {
	ErrorCode    int     `json:"error_code"`
	ErrorMessage string  `json:"error_message"`
	Body         *Author `json:"body"`
}

type GetAllAuthors struct {
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Body         []*Author `json:"body"`
}

type AuthorRes struct {
	ErrorCode    int     `json:"error_code"`
	ErrorMessage string  `json:"error_message"`
	Body         *Author `json:"body"`
}

type CreateAuthorReq struct {
	Name string `json:"name"`
}

type UpdateAuthorReq struct {
	Name string `json:"name"`
}

type DeleteAuthorRes struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
