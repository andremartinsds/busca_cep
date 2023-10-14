package resources

import (
	"busca-cep/src/services"
	"busca-cep/src/types"
	"encoding/json"
	"net/http"
)

func BuscaCepResource(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/cep" {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := req.URL.Query().Get("cep")
	if len(cepParam) > 8 || len(cepParam) < 8 {
		errorResponse := types.ErrorResponse{
			Error: "Size error",
		}
		resError, _ := json.Marshal(errorResponse)
		res.Write(resError)
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if cepParam == "" {
		errorResponse := types.ErrorResponse{
			Error: "cep vazio",
		}
		resErr, _ := json.Marshal(errorResponse)
		res.Write(resErr)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	cepResponse, err := services.BuscaCep(cepParam)

	if err != nil {
		errorResponse := types.ErrorResponse{
			Error: err.Error(),
		}
		resErr, _ := json.Marshal(errorResponse)
		res.Write(resErr)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	response, err := json.Marshal(types.Response{
		Data: *cepResponse,
	})

	if err != nil {
		errorResponse := types.ErrorResponse{
			Error: err.Error(),
		}
		resErr, _ := json.Marshal(errorResponse)
		res.Write(resErr)
		res.WriteHeader(http.StatusBadRequest)
	}

	res.Write(response)
	return
}
