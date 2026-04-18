# Imagen base
FROM golang:1.25

# Directorio de trabajo
WORKDIR /app

# Copiar go.mod primero (mejora cache)
COPY go.mod ./
RUN go mod tidy

# Copiar todo el código
COPY . .

# Compilar
RUN go build -o main ./cmd

# Ejecutar
CMD ["./main"]