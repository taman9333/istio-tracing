FROM golang:1.15

WORKDIR /app

RUN go get github.com/githubnemo/CompileDaemon

COPY go.mod ./

RUN go mod vendor

COPY . .

EXPOSE 3000

CMD ["CompileDaemon", "-command", "./xyz"]