package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
		}
		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo: %v\n", err)
		}

		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("Cep: %v, End: %v, Bairro: %v, Cidade: %v, UF: %v", data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.UF))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}
		fmt.Println("Arquivo criado com sucesso!")
		fmt.Fprintf(os.Stdout, "Cep: %v, End: %v, Bairro: %v, Cidade: %v, UF: %v", data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.UF)
	}
}
