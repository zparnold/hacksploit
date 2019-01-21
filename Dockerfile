# build stage
FROM golang AS build-env
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/zparnold/hacksploit/
COPY . /go/src/github.com/zparnold/hacksploit/
RUN dep ensure && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./src/...

# final stage
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-env /go/src/github.com/zparnold/hacksploit/app .
CMD ["./app"]