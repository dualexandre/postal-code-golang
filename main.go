package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type PostalCode struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", GetPostalCodeHandler)
	http.ListenAndServe(":8000", nil)
}

func GetPostalCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	codeParam := r.URL.Query().Get("code")
	if codeParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	code, err := GetPostalCode(codeParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(code)
}

func GetPostalCode(code string) (*PostalCode, error) {
	res, err := http.Get("https://viacep.com.br/ws/" + code + "/json/")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var c PostalCode
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
