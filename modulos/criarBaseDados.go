package modulos

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

/*Esta função cria uma base de dados a ser utilizada pela aplicação. */

func CreateBaseDados(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

/*Esta função cria a tabela a ser utilizada pela aplicação. Impossibilitado de usar as váriaveis do tipo "$1",
para montar a execução do comando. Deve ser estudada uma melhor forma de criar as tabelas, facilitando a criação
de várias tabelas com a passagem de parâmetro, com o nome da tabela. */

func CreateTabelaDadosClientes(db *sql.DB) error {
	schema := "DROP TABLE IF EXISTS dadosclientes; CREATE TABLE dadosclientes (id serial PRIMARY KEY, cpf VARCHAR (50) NOT NULL, private INTEGER, incompleto INTEGER, dataultcompra DATE, ticketmed NUMERIC (10, 2), ticketultcompra NUMERIC (10, 2), lojausual VARCHAR (50), lojaultcompra VARCHAR (50));"

	result, err := db.Exec(schema)

	if err != nil {
		log.Panic(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Panic(err)
		return err
	}

	strconv.FormatInt(rowsAffected, 10)
	return nil
}
