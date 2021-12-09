FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN go get github.com/go-delve/delve/cmd/dlv@latest

RUN go build -o ./bin/server main.go

CMD ["/dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/bin/server"]
