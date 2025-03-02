FROM debian:latest

RUN apt-get update -y && apt-get install -y \
    curl \
    git \
    build-essential \
    libdlib-dev \
    libblas-dev \
    liblapack-dev \
    libjpeg62-turbo-dev

# Устанавливаем последнюю версию Go
ENV GO_VERSION=1.24.0
RUN curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz
RUN apt-get install -y libhdf5-dev

# Обновляем переменные окружения
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOTOOLCHAIN=local 

WORKDIR /app
COPY . /app

RUN go mod download
RUN export CPATH="/usr/include/hdf5/serial/" && go build -v -o main main.go

ENTRYPOINT ["/app/main"]
EXPOSE 8080

