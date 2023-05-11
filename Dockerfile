FROM golang:1.20-alpine AS builder

WORKDIR $GOPATH/src/learn-github-flow

COPY ./ ./

RUN go mod tidy && go mod download

RUN CGO_ENABLED=0 go build -o bin/app ./main.go

FROM alpine:latest

RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

WORKDIR /

COPY --from=builder /go/src/learn-github-flow/bin/app ./

EXPOSE 8080

ENTRYPOINT ["./app"]