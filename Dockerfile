
# ETAPA 1: Compilar la aplicacion Go
FROM golang:1.26-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o servidor main.go

# ETAPA 2: Crear la imagen final
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/servidor .
COPY static ./static

EXPOSE 8080

CMD ["./servidor"]
