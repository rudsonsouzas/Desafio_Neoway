package modulos

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
)

/*Esta função apaga todos os registros na tabela com as informações dos clientes. */

func DeleteDadosClientes(db *sql.DB) (int64, error) {
	sql := "DELETE FROM dadosclientes"
	result, err := db.Exec(sql)

	if err != nil {
		log.Panic(err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Panic(err)
		return 0, err
	}

	strconv.FormatInt(rowsAffected, 10)
	return rowsAffected, nil
}

/*Esta função realiza a consulta na tabela com os dados dos clientes e retorna o resultado desta consulta. */

func GetDadosClientes(db *sql.DB) ([]*DadosCliente, error) {
	rows, err := db.Query("SELECT cpf, private, incompleto, dataultcompra, ticketmed, ticketultcompra, lojausual, lojaultcompra FROM dadosclientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dadosCliente := make([]*DadosCliente, 0)
	for rows.Next() {
		dc := new(DadosCliente)
		err := rows.Scan(&dc.Cpf, &dc.Private, &dc.Incompleto, &dc.DataUltCompra, &dc.TicketMed, &dc.TicketUltCompra, &dc.LojaUsual, &dc.LojaUltCompra)
		if err != nil {
			return nil, err
		}

		dadosCliente = append(dadosCliente, dc)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return dadosCliente, nil
}

/*Esta função insere os dados na tabela com as informações dos clientes. */

func InsertDadosClientes(db *sql.DB, dc InsertDadosCliente) (int64, error) {
	sql := "INSERT INTO dadosclientes (cpf, private, incompleto, dataultcompra, ticketmed, ticketultcompra, lojausual, lojaultcompra) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) "

	cpf := NewNullString(dc.Cpf)
	private := NewNullInt(dc.Private)
	incompleto := NewNullInt(dc.Incompleto)
	dataUltCompra := NewNullString(dc.DataUltCompra)
	ticketMed := NewNullFloat(dc.TicketMed)
	ticketUltCompra := NewNullFloat(dc.TicketUltCompra)
	lojaUsual := NewNullString(dc.LojaUsual)
	lojaUltCompra := NewNullString(dc.LojaUltCompra)

	result, err := db.Exec(sql, cpf, private, incompleto, dataUltCompra, ticketMed, ticketUltCompra, lojaUsual, lojaUltCompra)

	if err != nil {
		log.Panic(err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Panic(err)
		return 0, err
	}

	strconv.FormatInt(rowsAffected, 10)
	return rowsAffected, nil
}

/*Esta função valida se a string repassada é nula, provendo o tratamento adequado. */

func NewNullString(s string) sql.NullString {
	if len(s) == 0 || s == "NULL" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

/*Esta função valida se o inteiro repassado é nulo, provendo o tratamento adequado. */

func NewNullInt(s string) sql.NullInt64 {
	if len(s) == 0 || s == "NULL" {
		return sql.NullInt64{}
	}

	i64, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return sql.NullInt64{
		Int64: i64,
		Valid: true,
	}
}

/*Esta função valida se o float repassado é nulo, provendo o tratamento adequado. */

func NewNullFloat(s string) sql.NullFloat64 {
	if len(s) == 0 || s == "NULL" {
		return sql.NullFloat64{}
	}

	sr := strings.Replace(s, ",", ".", -1)
	f64, err := strconv.ParseFloat(sr, 64)

	if err != nil {
		log.Fatal(err)
	}

	return sql.NullFloat64{
		Float64: f64,
		Valid:   true,
	}
}

/*Formaliza que os atributos do tipo insertDadosCliente serão strings a fim de facilitar a leitura das informações
presentes no arquivo csv/txt. */

type InsertDadosCliente struct {
	Cpf             string
	Private         string
	Incompleto      string
	DataUltCompra   string
	TicketMed       string
	TicketUltCompra string
	LojaUsual       string
	LojaUltCompra   string
}

/*Formaliza os tipos para os atributos do tipo DadosCliente, além de identar sua descrição no arquivo JSON. Essas
informações serão utilizadas para popular/manipular a tabela com os dados dos clientes. */

type DadosCliente struct {
	Cpf             NullString  `json:"Cpf"`
	Private         NullInt64   `json:"Private"`
	Incompleto      NullInt64   `json:"Incompleto"`
	DataUltCompra   NullString  `json:"DataUltCompra"`
	TicketMed       NullFloat64 `json:"TicketMed"`
	TicketUltCompra NullFloat64 `json:"TicketUltCompra"`
	LojaUsual       NullString  `json:"LojaUsual"`
	LojaUltCompra   NullString  `json:"LojaUltCompra"`
}
