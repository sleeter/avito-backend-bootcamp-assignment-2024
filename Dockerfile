FROM golang:1.22-alpine

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN mkdir -p /usr/local/bin/
RUN go build -v -o /usr/local/bin/app ./cmd/main.go

CMD ["app"]