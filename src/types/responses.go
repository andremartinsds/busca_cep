package types

type Response struct {
	Data ViaCep `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
