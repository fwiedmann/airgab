FROM golang:1.11.5 AS builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/fwiedmann/airgab/

COPY . .

RUN go get
RUN CGO_ENABLED=0 go build -a -o airgab .

FROM alpine:3.9

WORKDIR /airgab

RUN apk --no-cache add rsync openssh
RUN mkdir -pv /root/.ssh/ \
    && chmod 0700 /root/.ssh

COPY --from=builder /go/src/github.com/fwiedmann/airgab/airgab ./airgab

ENTRYPOINT [ "./airgab" ]