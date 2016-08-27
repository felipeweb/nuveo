package parser

import (
	"encoding/csv"
	"encoding/json"
	"github.com/felipeweb/nuveo/models"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func ToJSON(data io.Reader) ([]models.Client, error) {
	var client models.Client
	var pmap map[string]interface{}

	bytes, err := ioutil.ReadAll(data)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &pmap)

	if err != nil {
		return nil, err
	}

	for key, val := range pmap {
		switch strings.ToLower(key) {
		case "nome":
			client.Nome = val.(string)
		case "email":
			client.Email = val.(string)
		case "sexo":
			client.Sexo = val.(string)
		case "idade":
			client.Idade = val.(int)
		default:
			client.Outros[key] = val
		}
	}

	var clients []models.Client
	return append(clients, client), err
}

func ToCSV(data io.Reader) ([]models.Client, error) {
	reader := csv.NewReader(data)
	matriz, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var clients []models.Client

	for _, line := range matriz {
		var client models.Client

		for col, value := range line {
			switch col {
			case 0:
				client.Nome = value
			case 1:
				client.Email = value
			case 2:
				client.Sexo = value
			case 3:
				idade, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
				client.Idade = idade
			default:
				client.Outros[col] = value
			}
		}

		clients = append(clients, client)
	}

	return clients, err
}
