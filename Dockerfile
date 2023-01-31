FROM golang:latest

WORKDIR /compiler

COPY ./ ./

RUN go build -o interpreter

FROM alpine:latest

ENTRYPOINT ["./interpreter"]

CMD ["./interpreter","--help"]