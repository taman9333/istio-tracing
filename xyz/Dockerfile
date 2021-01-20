FROM golang:1.15

WORKDIR /app

COPY . .

RUN go get github.com/githubnemo/CompileDaemon

RUN go mod vendor

EXPOSE 3000

CMD ["CompileDaemon", "-command", "./xyz"]