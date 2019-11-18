FROM golang as builder

WORKDIR /go/src/app
COPY . .

RUN go build -ldflags "-linkmode external -extldflags -static" -a main.go

FROM scratch

WORKDIR /app
COPY --from=builder /go/src/app .

CMD ["./main"]