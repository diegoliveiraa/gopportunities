FROM golang:1.24-alpine AS builder

WORKDIR /app

# Instala ferramentas de build e dependências do SQLite
RUN apk add --no-cache make git gcc musl-dev

# Ativa o CGO
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
ENV PATH="/go/bin:$PATH"

# Instala o swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Executa build que gera binário com CGO
RUN make docker

# Imagem final mínima
FROM alpine:latest

WORKDIR /app

# Copia a libc (necessária para binário CGO)
RUN apk add --no-cache libc6-compat

# Copia binário gerado
COPY --from=builder /app/gopportunities .

EXPOSE 8080

CMD ["./gopportunities"]
