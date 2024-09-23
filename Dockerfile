FROM golang:1.21.6

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o tasks-app ./cmd/main.go

CMD ["./tasks-app"]