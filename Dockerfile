# Usando a imagem base do Golang
FROM golang:1.22

# Definindo o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiando os arquivos de configuração do Go (mod e sum)
COPY go.mod go.sum ./

# Instalando as dependências
RUN go mod download

# Copiando o arquivo .env
COPY .env ./

# Instalando a ferramenta de live reloading air
RUN go install github.com/cosmtrek/air@latest

# Copiando o restante do código da aplicação
COPY . .

# Comando para iniciar a aplicação usando air
CMD ["air"]
