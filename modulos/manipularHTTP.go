package modulos

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func (ci *ContextInjector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ci.h.ServeHTTP(w, r.WithContext(ci.ctx))
}

/*Esta função remove todos os registros da tabela onde estão os dados dos clientes. */

func RemoveDadosClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	db, ok := r.Context().Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "could not get database connection pool from context", 500)
		return
	}

	rowsDel, err := DeleteDadosClientes(db)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	var tpld = template.Must(template.ParseGlob("templates/*"))
	tpld.ExecuteTemplate(w, "Delete", rowsDel)
	//fmt.Fprintf(w, "%s", rowsAffected)
}

/*Esta função popula as informações do arquivo do tipo txt/csv na tabela criada no banco de dados. */

func PopulateDadosClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	db, ok := r.Context().Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "could not get database connection pool from context", 500)
		return
	}

	/*Remove todos os registros da tabela com os dados dos clientes, para posterior inclusão de novas informações. */
	_, err1 := DeleteDadosClientes(db)
	if err1 != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	/*Busca dados de arquivo do tipo csv/txt, salvo na pasta arquivos. */
	dadosClientes := GetArquivo()

	/*Insere os registros na tabela. */
	for _, dc := range dadosClientes {
		valid := (IsCNPJ(dc.Cpf) || IsCPF(dc.Cpf))
		if !valid {
			fmt.Println("CNPJ/CPF inválido: " + dc.Cpf)
			continue
		}

		Clean(&dc.Cpf)
		Clean(&dc.LojaUsual)
		Clean(&dc.LojaUltCompra)

		_, err2 := InsertDadosClientes(db, dc)

		if err2 != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

	}

	/*sql := "SELECT COUNT(distinct id) FROM dadosclientes;"

	result, err := db.Exec(sql)

	if err != nil {
		log.Panic(err)
		return
	}*/

	var tplm = template.Must(template.ParseGlob("templates/*"))
	tplm.ExecuteTemplate(w, "Migrar", nil)
}

/*Esta função exibe as informações da tabela com os dados do cliente no formato JSON. */

func ShowDadosClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	db, ok := r.Context().Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "could not get database connection pool from context", 500)
		return
	}

	dadosClientes, err := GetDadosClientes(db)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	/*notjs, err := UnmarshalJSON(dadosClientes)
	if err != nil {
		return
	}*/

	var tpls = template.Must(template.ParseGlob("templates/*"))
	tpls.ExecuteTemplate(w, "Show", dadosClientes)
}

/*Esta função exibe as informações da tabela com os dados do cliente no formato JSON. */

func ShowDadosClientesJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	db, ok := r.Context().Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "could not get database connection pool from context", 500)
		return
	}

	dadosClientes, err := GetDadosClientes(db)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "%s", "[")

	size := (len(dadosClientes) - 1)
	for i, dc := range dadosClientes {
		js, err := json.Marshal(dc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if size == i {
			fmt.Fprintf(w, "%s", js)
		} else {
			fmt.Fprintf(w, "%s,", js)
		}

	}
	fmt.Fprintf(w, "%s", "]")
}

/*Esta função exibe as informações da Home. */

func ShowHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var tplh = template.Must(template.ParseGlob("templates/*"))
	tplh.ExecuteTemplate(w, "Home", nil)

}

/*Formaliza o tipo ContextInjector e os tipos de seus atributos. */

type ContextInjector struct {
	ctx context.Context
	h   http.Handler
}
