
FROM golang:latest

WORKDIR /app 

COPY go.mod ./
RUN go mod download 

COPY ./ ./

RUN go build -o /go-docker-demo 

CMD [ "/go-docker-demo" ]