package models

import (
	"alura/webapp/db"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id    int
	Nome  string
	Desc  string
	Preco float64
	Quant int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectTodosProdutos, err := db.Query("select * from produtos order by id")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectTodosProdutos.Next() {
		var id int
		var nome, descricao string
		var preco float64
		var quant int
		err = selectTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quant)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Desc = descricao
		p.Preco = preco
		p.Quant = quant

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func NovoProduto(nome, descricao string, preco float64, quant int) {
	db := db.ConectaComBancoDeDados()

	insert, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(nome, descricao, preco, quant)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()

	deletaProduto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}
	deletaProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id = $1", id)
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
		produtoParaAtualizar.Desc = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quant = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, desc string, preco float64, quant int) {
	db := db.ConectaComBancoDeDados()

	atualizaProduto, err := db.Prepare("update produtos set nome = $1, descricao = $2, preco = $3, quantidade = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(nome, desc, preco, quant, id)
	defer db.Close()
}
