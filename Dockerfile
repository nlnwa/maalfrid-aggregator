FROM golang:alpine

RUN apk add --no-cache --update alpine-sdk protobuf protobuf-dev

COPY . /go/src/github.com/nlnwa/maalfrid-language-detector

RUN cd /go/src/github.com/nlnwa/maalfrid-aggregator \
&& go generate github.com/nlnwa/maalfrid-aggregator/api \
&& go get github.com/golang/dep/cmd/dep \
&& dep ensure -vendor-only \
&& VERSION=$(./scripts/git-version) \
CGO_ENABLED=0 \
go install -a -tags netgo -v -ldflags "-w -X github.com/nlnwa/maalfrid-aggregator/version.Version=$(VERSION)" \
github.com/nlnwa/maalfrid-aggregator/cmd/...
# -w Omit the DWARF symbol table.
# -X Set the value of the string variable in importpath named name to value.

FROM scratch
LABEL maintainer="marius.beck@nb.no"

COPY --from=0 /go/bin/maalfrid /

ENV PORT=8672

ENTRYPOINT ["/maalfrid-aggregator"]

EXPOSE 8672
