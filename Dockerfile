######## Start the build stage #######
FROM golang:1.14 as builder

LABEL maintainer="Dmytro Nakonechnyy <https://github.com/nakonechniyd>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


######## Start the relese stage #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main", "-h", "0.0.0.0", "-p", "8080"]
