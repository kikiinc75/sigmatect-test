FROM golang:1.18-alpine

#set working directory
WORKDIR /go/src/sigmatech-test

# copy golang module dependencies
COPY go.mod go.mod
COPY go.sum go.sum

# download all golang module dependencies
RUN go mod download

COPY . .

RUN go build -o /go/bin/sigmatech-test main.go

CMD [ "/go/bin/sigmatech-test" ]