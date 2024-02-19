FROM golang:latest

WORKDIR /cmd/checkwpjson

COPY . .
RUN go build -o checkwpjson ./\cmd/\checkwpjson/\checkwpjson.go
ENTRYPOINT [ "./checkwpjson" ]

