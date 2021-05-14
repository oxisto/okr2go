FROM debian:buster-slim

ARG TARGETOS
ARG TARGETARCH

ADD build/okr2go-${TARGETOS}-${TARGETARCH}.tar.gz .
RUN ls *

ENTRYPOINT [ "okr2go" ]
