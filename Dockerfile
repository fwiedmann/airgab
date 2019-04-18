FROM alpine:3.9

LABEL maintainer=https://github.com/fwiedmann github-project=https://github.com/fwiedmann/airgab

RUN apk --no-cache add rsync openssh bash curl

RUN addgroup pilot \
    && adduser -D pilot -G pilot

RUN mkdir -pv /home/pilot/.ssh/ \
    && chmod 0700 /home/pilot/.

WORKDIR /home/pilot

COPY entrypoint.sh ./entrypoint.sh
COPY airgab /usr/local/bin/airgab

ENTRYPOINT [ "./entrypoint.sh" ]
HEALTHCHECK --interval=10s --timeout=5s --retries=3 CMD curl -f http://localhost:9100/metrics || exit 1