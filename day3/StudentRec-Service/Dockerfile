FROM golang:1.16

WORKDIR /app
COPY . .
EXPOSE 8080

RUN go mod tidy

RUN go get github.com/codegangsta/gin

CMD ["gin", "-p", "8080", "run", "main.go"]