package api

type User struct {
	ID   uint8  `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
	Status  uint8  `json:"status"`
}
