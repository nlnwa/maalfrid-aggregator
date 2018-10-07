FROM golang:alpine

RUN apk add --no-cache --update alpine-sdk protobuf protobuf-dev

COPY . /go/src/github.com/nlnwa/maalfrid-aggregator

RUN cd /go/src/github.com/nlnwa/maalfrid-aggregator \
&& go generate github.com/nlnwa/maalfrid-aggregator/maalfrid/aggregator \
# && go get github.com/golang/dep/cmd/dep \
# && dep ensure -vendor-only \
&& go get ./... \
&& VERSION=$(./scripts/git-version) \
CGO_ENABLED=0 \
go install -a -tags netgo -v -ldflags "-w -X github.com/nlnwa/maalfrid-aggregator/version.Version=$(VERSION)" \
github.com/nlnwa/maalfrid-aggregator/cmd/...
# -w Omit the DWARF symbol table.
# -X Set the value of the string variable in importpath named name to value.

FROM scratch
LABEL maintainer="marius.beck@nb.no"

COPY --from=0 /go/bin/maalfrid-aggregator /

ENV PORT=8672 \
    DB_HOST=localhost \
    DB_PORT=28015 \
    DB_NAME=test \
    DB_USER=admin \
    DB_PASSWORD=""

ENTRYPOINT ["/maalfrid-aggregator"]

EXPOSE 8672
