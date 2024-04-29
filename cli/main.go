package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	for _, code := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + code + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v \n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v \n", err)
		}

		var data PostalCode
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v \n", err)
		}

		file, err := os.Create("address.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v \n", err)
		}
		defer file.Close()

		_, err = file.WriteString(
			fmt.Sprintf("Postal Code: %s, Street: %s, State: %s. \n", data.Cep, data.Logradouro, data.Uf),
		)

		fmt.Println(data)
	}
}
