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
	selectDeTodosOsProdutos, err := dbCon.Query("select * from produtos order by id asc")
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

		p.Id = id
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
	defer func(dbCon *sql.DB) {
		err := dbCon.Close()
		if err != nil {
			panic(err.Error())
		}
	}(dbCon)
}

func DeletaProduto(id string) {
	dbCon := db.ConectaComBancoDeDados()

	deletarProduto, err := dbCon.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(id)

	err = dbCon.Close()
	if err != nil {
		return
	}
}

func EditaProduto(id string) Produto {
	dbCon := db.ConectaComBancoDeDados()
	produtoDoBanco, err := dbCon.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer func(dbCon *sql.DB) {
		err := dbCon.Close()
		if err != nil {
			panic(err.Error())
		}
	}(dbCon)
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	dbCon := db.ConectaComBancoDeDados()

	AtualizaProduto, err := dbCon.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)

	err = dbCon.Close()
	if err != nil {
		return
	}
}
