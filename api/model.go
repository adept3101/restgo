package api

type User struct {
	ID   uint8  `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type UsersResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    []User `json:"data"`
}

type UserResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    User   `json:"data"`
}
