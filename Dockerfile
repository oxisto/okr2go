FROM debian:buster-slim

ARG TARGETOS
ARG TARGETARCH

ADD build/okr2go-${TARGETOS}-${TARGETARCH}.tar.gz .

EXPOSE 4300

ENTRYPOINT ["/okr2go"]
