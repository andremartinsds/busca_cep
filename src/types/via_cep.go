package types

type ViaCep struct {
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
}
