# Desafio_Neoway

## Objetivo
Este repositório tem como objetivo disponibilizar os arquivos do aplicativo desenvolvido para migração das informações do arquivo 'base_teste.txt' para uma base de dados do Postgres.

O desenvolvimento do aplicativo foi realizado na linguagem Golang e virtualizado através do Docker.

## Requisitos
Para execução deste aplicativo é necessário a instalação do Docker e do Docker Compose.

Utilizamos para desenvolvimento e testes a versão 1.13.1.

## Ações para execução da aplicação 

Primeiro, baixe os arquivos do repositório para seu ambiente de trabalho.

Em seguida execute o script 'iniciar.sh', presente na raiz do repositório, para realizar a inicialização dos containers e da aplicação.

Acesse em seu navegador de preferência o seguinte endereço: http://localhost:8090

Será exibida a tela do link: https://drive.google.com/file/d/1foLhz9Ms1aXlKTGopw5DpOpgVTGnztkX/view?usp=sharing

A aplicação possui quatro botões: 
• "Dados do Cliente", exibe as informações dos clientes no formato HTML;
• "Dados do Cliente - JSON", exibe as informações dos clientes no formato JSON;
• "Excluir Dados", remove todos os dados da aplicação;
• "Migrar Dados", migra todos os dados do arquivo txt/csv para a aplicação.

Para encerrar a aplicação, execute o script 'encerrar.sh', presente na raiz do repositório.

## Autor
Rudson Vieira de Souza
