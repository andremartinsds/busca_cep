package types

// apenas para testes, o ideal é struct diferentes para cada resposta
type TypeCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"-"`
	Uf          string `json:"-"`
	Ibge        string `json:"-"`
	Gia         string `json:"-"`
	Ddd         string `json:"-"`
	Siafi       string `json:"-"`

	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}
