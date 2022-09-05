package dto

type RegisterUserRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	CellNumber   string `json:"cell_number"`
	NationalCode string `json:"national_code"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

type RegisterUserResponse struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	CellNumber   string `json:"cell_number"`
	NationalCode string `json:"national_code"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	CellNumber   string `json:"cell_number"`
	NationalCode string `json:"national_code"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
}
