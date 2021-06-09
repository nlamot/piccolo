FROM golang:1.16
WORKDIR /go/src/github.com/nlamot/piccolo/
COPY . .
RUN make piccolo

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/nlamot/piccolo/bin/piccolo .
CMD ["./piccolo"]  