FROM golang:1.12-alpine as builder

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

ENV GO111MODULE=on

WORKDIR /go/src/api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o server .

FROM scratch

COPY --from=builder /go/src/api/server /api/
EXPOSE 8888
WORKDIR /api
CMD ["./server"]
