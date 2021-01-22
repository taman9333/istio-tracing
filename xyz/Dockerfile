FROM golang:1.15

WORKDIR /app

RUN go get github.com/codegangsta/gin

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN go mod vendor

EXPOSE 3000

CMD ["gin"]