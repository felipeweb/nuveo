package models

type Client struct {
	Nome   string
	Email  string
	Sexo   string
	Idade  string
	Outros map[interface{}]interface{}
}
