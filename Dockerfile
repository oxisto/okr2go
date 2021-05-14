FROM debian:buster-slim

ADD okr2go .

ARG TARGETPLATFORM
ARG TARGETARCH
ARG TARGETOS
RUN echo "I want okr2go-${TARGETOS}-${TARGETARCH}.zip"

ENTRYPOINT [ "okr2go" ]
