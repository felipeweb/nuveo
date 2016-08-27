package parser

import (
	"encoding/csv"
	"encoding/json"
	"github.com/felipeweb/nuveo/models"
	"io"
	"io/ioutil"
	"strings"
	"github.com/felipeweb/gopher-utils"
)

func ToJSON(data io.Reader) ([]models.Client, error) {
	client := models.Client{Outros:map[interface{}]interface{}{}}
	var clientMap map[string]interface{}

	bytes, err := ioutil.ReadAll(data)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &clientMap)

	if err != nil {
		return nil, err
	}

	for key, value := range clientMap {
		switch strings.ToLower(key) {
		case "nome":
			client.Nome = strings.TrimSpace(gopher_utils.ToStr(value))
		case "email":
			client.Email = strings.TrimSpace(gopher_utils.ToStr(value))
		case "sexo":
			client.Sexo = strings.TrimSpace(gopher_utils.ToStr(value))
		case "idade":
			client.Idade = strings.TrimSpace(gopher_utils.ToStr(value))
		default:
			client.Outros[strings.ToLower(key)] = strings.TrimSpace(gopher_utils.ToStr(value))
		}
	}

	var clients []models.Client
	clients = append(clients, client)
	return clients, err
}

func ToCSV(data io.Reader) ([]models.Client, error) {
	reader := csv.NewReader(data)
	matriz, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var clients []models.Client

	for line, types := range matriz {
		if line != 0 {
			client := models.Client{Outros:map[interface{}]interface{}{}}
			for col, value := range types {
				switch col {
				case 0:
					client.Nome = strings.TrimSpace(value)
				case 1:
					client.Email = strings.TrimSpace(value)
				case 2:
					client.Sexo = strings.TrimSpace(value)
				case 3:
					client.Idade = strings.TrimSpace(value)
				default:
					colName := matriz[0][col]
					colName = strings.TrimSpace(colName)
					client.Outros[strings.ToLower(colName)] = strings.TrimSpace(value)
				}
			}

			clients = append(clients, client)
		}
	}

	return clients, err
}
