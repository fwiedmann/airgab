FROM golang:1.11.5 AS builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/fwiedmann/airgab/

COPY . .

RUN go get
RUN CGO_ENABLED=0 go build -a -o airgab .

FROM alpine:3.9

RUN apk --no-cache add rsync openssh bash curl

RUN addgroup pilot \
    && adduser -D pilot -G pilot

RUN mkdir -pv /home/pilot/.ssh/ \
    && chmod 0700 /home/pilot/.

WORKDIR /home/pilot

COPY entrypoint.sh ./entrypoint.sh
COPY --from=builder /go/src/github.com/fwiedmann/airgab/airgab ./airgab

ENTRYPOINT [ "./entrypoint.sh" ]
HEALTHCHECK --interval=10s --timeout=5s --retries=3 CMD curl -f http://localhost:9100/metrics || exit 1