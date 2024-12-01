FROM golang:1.23

WORKDIR /app

COPY go.mod .
COPY main.go .

#para que se instalen los paquetes
RUN go get


RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]