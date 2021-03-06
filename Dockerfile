FROM golang:1.16
WORKDIR $GOPATH/src/github.com/losglennos/metrics-appinsights
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["metrics-appinsights"]