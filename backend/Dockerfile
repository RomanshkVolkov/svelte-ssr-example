# Team: Romanshk Volkov - https://github.com/RomanshkVolkov
# Team: Diegode - https://github.com/diegode-tsx
# Team: Alexandergv2117 - https://github.com/Alexandergv2117
# Project: api

FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build -o /api ./cmd/

# runner light enviroment
FROM alpine:latest

WORKDIR /srv

COPY --from=builder /api .

EXPOSE 8080

CMD ["./api"]
