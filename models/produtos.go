package models

import (
	"GoLangAppWeb/db"
	"database/sql"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	dbCon := db.ConectaComBancoDeDados()
	selectDeTodosOsProdutos, err := dbCon.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id int
		var nome, descricao string
		var preco float64
		var quantidade int

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	defer func(dbCon *sql.DB) {
		err := dbCon.Close()
		if err != nil {
			panic(err.Error())
		}
	}(dbCon)
	return produtos
}
