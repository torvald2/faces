FROM golang:1.23

# Устанавливаем необходимые библиотеки и зависимости
RUN apt-get update -y && apt-get install -y \
    libdlib-dev \
    libblas-dev \
    liblapack-dev \
    libjpeg62-turbo-dev \
    libhdf5-dev \
    build-essential \
    git

WORKDIR /app
COPY . /app

# Загружаем зависимости и собираем проект
RUN go mod download
RUN go build -v -o main main.go

ENTRYPOINT ["/app/main"]
EXPOSE 8080


