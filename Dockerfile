FROM debian:latest

RUN apt-get update -y && apt-get install -y \
    curl \
    git \
    build-essential \
    libdlib-dev \
    libblas-dev \
    liblapack-dev \
    libjpeg62-turbo-dev

ENV GO_VERSION=1.21.5
RUN curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz

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
