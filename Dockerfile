FROM debian

ENV http_proxy=http://dockerhost:5555
ENV htts_proxy=http://dockerhost:5555


RUN  apt-get update -y
RUN  apt-get install software-properties-common -y

##RUN  add-apt-repository ppa:longsleep/golang-backports
##RUN  add-apt-repository ppa:kagamih/dlib
##RUN  apt-get update -y   

RUN apt-get   install golang-go -y
RUN apt-get install git -y

RUN  apt-get install libdlib-dev -y
RUN  apt-get install libblas-dev  -y
RUN  apt-get install liblapack-dev  -y
RUN  apt-get install libjpeg62-turbo-dev -y

ENV GOOS=linux
ENV GOARCH=amd64 

RUN mkdir app
COPY . /app

WORKDIR /app

RUN go mod download
RUN export CPATH="/usr/include/hdf5/serial/"
RUN go build -v main.go

ENTRYPOINT /app/main
EXPOSE 8080








