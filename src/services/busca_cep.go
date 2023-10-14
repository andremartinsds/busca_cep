package services

import (
	"busca-cep/src/types"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func BuscaCep(cep string) (*types.ViaCep, error) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	data_read, error_read_all := io.ReadAll(resp.Body)

	if error_read_all != nil {
		return nil, error_read_all
	}

	defer resp.Body.Close()

	var via_cep types.ViaCep

	json.Unmarshal(data_read, &via_cep)

	if via_cep == (types.ViaCep{}) {
		return nil, errors.New("nada foi encontrado")
	}

	return &via_cep, nil

}
