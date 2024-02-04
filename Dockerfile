FROM debian:oldstable


##RUN  echo "deb http://ftp.uk.debian.org/debian/ unstable main" |  tee -a /etc/apt/sources.list
RUN  apt-get update -y
RUN  apt-get install software-properties-common -y

#RUN  add-apt-repository ppa:/gophers/archive
##RUN  apt-update -y
##RUN  add-apt-repository ppa:kagamih/dlib
##RUN  apt-get update -y   

#RUN apt-get  -t buster-backports install golang-1.14-go -y
#RUN apt-get install golang-1.15-go -y
RUN apt-get install curl -y
RUN apt-get install git -y



RUN apt install golang


RUN  apt-get install libdlib-dev -y
RUN  apt-get install libblas-dev  -y
RUN  apt-get install liblapack-dev  -y
RUN  apt-get install libjpeg62-turbo-dev -y

ENV GOOS=linux
ENV GOARCH=amd64 

RUN mkdir app
COPY . /app

WORKDIR /app
RUN apt-get install build-essential -y

ENV PATH=$PATH:/usr/local/go/bin
RUN  go get -u gonum.org/v1/gonum
RUN go mod download
RUN export CPATH="/usr/include/hdf5/serial/"
RUN go build -v main.go

ENTRYPOINT /app/main
EXPOSE 8080








