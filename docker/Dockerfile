FROM golang:1.19 as builder

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd               ./cmd

RUN ls /app/
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -o meta ./cmd/api

FROM scratch
WORKDIR /app
COPY --from=builder /app/meta /app/

EXPOSE 8000
ENTRYPOINT ["/app/meta"]