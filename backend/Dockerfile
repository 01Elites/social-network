FROM golang:latest

EXPOSE 8081

COPY . /social-backend

WORKDIR /social-backend

RUN go build -o socialNetwork cmd/socialNetwork/main.go

CMD ["./socialNetwork"]