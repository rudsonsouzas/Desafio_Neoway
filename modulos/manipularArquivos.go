package modulos

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

/*Esta função usa o pacote OS para abrir o arquivo de dados a importar, repassando-o a um scanner, do pacote bufio,
que extrairá cada linha do arquivo e a passará pela função quebraLinhasParaColunas, que dividará cada trecho da linha
para a coluna correta a ser inserida no banco de dados.
*/

func GetArquivo() []InsertDadosCliente {
	file, err := os.Open("arquivos/base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	dadosCliente := make([]InsertDadosCliente, 0)

	for scanner.Scan() {
		dadosCliente = append(dadosCliente, QuebraLinhasParaColunas(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ao ler a linha encontramos este erro:", err)
	}

	return dadosCliente
}

/* Esta função utiliza expressões regulares para dividir cada linha repassada, nas colunas a serem gravadas no banco de
dados. As expressões regulares foram utilizadas pois não havia nenhuma formatação ou identificador que possibilitasse
a divisão correta das linhas em colunas.*/

func QuebraLinhasParaColunas(scannerText string) InsertDadosCliente {
	r := regexp.MustCompile("[^\\s]+")
	splitLine := r.FindAllString(scannerText, -1)

	return InsertDadosCliente{
		Cpf:             splitLine[0],
		Private:         splitLine[1],
		Incompleto:      splitLine[2],
		DataUltCompra:   splitLine[3],
		TicketMed:       splitLine[4],
		TicketUltCompra: splitLine[5],
		LojaUsual:       splitLine[6],
		LojaUltCompra:   splitLine[7],
	}
}
