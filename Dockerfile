FROM debian:buster-slim

ARG TARGETOS
ARG TARGETARCH

ADD okr2go .

ADD okr2go-${TARGETOS}-${TARGETARCH}.tar.gz .
RUN ls *

ENTRYPOINT [ "okr2go" ]
