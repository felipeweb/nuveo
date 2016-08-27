package models

type Client struct {
	Nome   string
	Email  string
	Sexo   string
	Idade  int
	Outros map[interface{}]interface{}
}
