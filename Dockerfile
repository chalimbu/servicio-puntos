FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build cmd/main.go

EXPOSE 4000

CMD [ "./main" ]