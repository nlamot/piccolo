FROM vektra/mockery:v2 as mockery

FROM golang:1.16

COPY --from=mockery /usr/local/bin/mockery /bin/mockery