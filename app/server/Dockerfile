FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

COPY ./migrations ./migrations

COPY Makefile ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/server

EXPOSE 1323

CMD ["/bin/server"]
