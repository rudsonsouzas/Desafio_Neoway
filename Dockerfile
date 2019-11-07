FROM golang

# Seta o diretório de trabalho 
WORKDIR $GOPATH/src/desafio_neoway

COPY . .

# Obtêm os drivers de conexão com o Postgres neste Git
RUN go get github.com/lib/pq
RUN go get -d -v ./...

# Instala os drivers
RUN go install -v ./...

# Parâmetro de configuração da porta de conexão da aplicação
EXPOSE 8090 
