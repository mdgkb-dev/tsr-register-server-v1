FROM golang:alpine

WORKDIR /app

RUN apk --no-cache add alpine-sdk

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go install github.com/cespare/reflex@latest

CMD [ "make", "run" ]
