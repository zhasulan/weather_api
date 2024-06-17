FROM golang:1.22-alpine as builder

WORKDIR /src

COPY ./ .

RUN unset GOPATH && GOOS=linux GOARCH=386 go mod tidy && go build -v -o ./build/run main.go

FROM alpine
RUN apk --no-cache add tzdata
WORKDIR /app

COPY --from=builder /src/config/conf.json /app/config/conf.json
COPY --from=builder /src/build/run .

CMD ["/app/run"]