package main

import (
	"context"
	"desafio_neoway/modulos"
	"log"
	"net/http"
)

func (ci *ContextInjector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ci.h.ServeHTTP(w, r.WithContext(ci.ctx))
}

/*Função principal do programa que cria a base de dados e os handles http para visualização dos dados. */

func main() {
	db, err := modulos.CreateBaseDados("postgres://postgres:@P0st123@db/neoway?sslmode=disable")

	if err != nil {
		log.Panic(err)
	}

	ctx := context.WithValue(context.Background(), "db", db)

	err1 := modulos.CreateTabelaDadosClientes(db)

	if err1 != nil {
		log.Panic(err)
	}

	http.Handle("/", &ContextInjector{ctx, http.HandlerFunc(modulos.ShowHome)})
	http.Handle("/delete", &ContextInjector{ctx, http.HandlerFunc(modulos.RemoveDadosClientes)})
	http.Handle("/show", &ContextInjector{ctx, http.HandlerFunc(modulos.ShowDadosClientes)})
	http.Handle("/showjson", &ContextInjector{ctx, http.HandlerFunc(modulos.ShowDadosClientesJSON)})
	http.Handle("/migrar", &ContextInjector{ctx, http.HandlerFunc(modulos.PopulateDadosClientes)})
	http.ListenAndServe(":8090", nil)
}

/*Formaliza o tipo ContextInjector e os tipos de seus atributos. */

type ContextInjector struct {
	ctx context.Context
	h   http.Handler
}
