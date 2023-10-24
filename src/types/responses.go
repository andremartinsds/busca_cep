package types

type Response struct {
	Data TypeCep `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
