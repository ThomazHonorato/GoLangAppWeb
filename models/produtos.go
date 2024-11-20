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

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	dbCon := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := dbCon.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer dbCon.Close()
}
