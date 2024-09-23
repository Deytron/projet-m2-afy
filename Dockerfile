FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY assets assets
COPY html html
COPY  *.go .env ./

EXPOSE 8081

RUN CGO_ENABLED=1 GOOS=linux go build -o /afy

CMD ["/afy"]