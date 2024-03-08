FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /weblog

ENV PORT=8124
EXPOSE $PORT

CMD ["/weblog", "serve"]
