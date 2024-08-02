FROM golang:1.22-alpine3.20

WORKDIR /app

COPY main.go .

RUN go mod init sample
RUN go mod tidy

CMD ["go","run","main.go"]