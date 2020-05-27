FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# app dependences
#COPY ./data.zip /tmp/data/data.zip
#COPY ./dataRate.zip /tmp/data/data.zip
RUN apt-get update
RUN apt-get -y install unzip
#RUN unzip ./data.zip -d /base/
#RUN unzip ./dataRate.zip -d /base/
#RUN unzip /tmp/data/data.zip -d /base/

#VOLUME ["./"]

EXPOSE 80
#RUN go build .
CMD ["app", ":80"]
#CMD ["go run .", ":80"]