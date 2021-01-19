FROM golang:1.15

WORKDIR /app

COPY . .

RUN go mod vendor

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 3000

CMD ["CompileDaemon", "-command", "./xyz"]