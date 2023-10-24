package services

import (
	"busca-cep/src/types"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

func BuscaCep(cep string, req *http.Request) (*types.TypeCep, error) {

	resp, err := ConcorrenciaApis(cep, req)

	if err != nil {
		return nil, err
	}

	dataResponse, errorResponse := io.ReadAll(resp.Body)

	if errorResponse != nil {
		return nil, errorResponse
	}

	defer resp.Body.Close()

	var typeCep types.TypeCep

	json.Unmarshal(dataResponse, &typeCep)

	if typeCep == (types.TypeCep{}) {
		return nil, errors.New("nada foi encontrado")
	}

	return &typeCep, nil

}

func ConcorrenciaApis(cep string, req *http.Request) (*http.Response, error) {
	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	urlViaCep := "https://viacep.com.br/ws/" + cep + "/json/"
	urlApiCep := "https://cdn.apicep.com/file/apicep/" + cep[:5] + "-" + cep[5:] + ".json"

	log.Println(urlApiCep)

	respChan := make(chan *http.Response)
	errChan := make(chan error)

	go func() {
		log.Println("executando api cep")
		resp, err := http.Get(urlApiCep)
		if err != nil {
			errChan <- err
		}
		respChan <- resp
		cancel()
	}()

	go func() {
		time.Sleep(5 * time.Second) // para testes
		log.Println("executando via cep")
		select {
		case <-ctx.Done():
			log.Print("primeira resposta já atendida")
			return
		default:
			resp, err := http.Get(urlViaCep)
			if err != nil {
				errChan <- err
			}

			respChan <- resp
			cancel()
		}
	}()

	select {
	case resp := <-respChan:
		log.Println("Resposta processada")
		return resp, nil
	case err := <-errChan:
		log.Print("houve um erro no processamento")
		return nil, err
	}
}
