################################################################################
# STAGE 1a: Build master binaries                                              #
################################################################################

FROM golang:1.12-buster as gobuilder
ENV GOPROXY="https://proxy.golang.org"

RUN set -e && \
  apt update && \
  apt install -y protobuf-compiler && \
  mkdir /root/qmstr/

WORKDIR /root/qmstr

COPY ./ ./

RUN set -e && \
  go build -o ./out/fasten-reporter-standalone ./modules/reporters/fasten-reporter-standalone

################################################################################
# STAGE 2: Build master container image (deploy)                               #
################################################################################

FROM python:3.6-slim-buster as deploy

# Required QMSTR directories
ENV QMSTR_DIRS="/var/qmstr/ /var/cache/qmstr/ /var/lib/qmstr/"

# Copy binaries from build stage
COPY --from=gobuilder /root/qmstr/out/* /usr/local/bin/

RUN set -e && \
  addgroup qmstr && \
  adduser --system qmstr --ingroup qmstr && \
  mkdir -p ${QMSTR_DIRS} && \
  chown -R qmstr ${QMSTR_DIRS}

WORKDIR /home/qmstr
USER qmstr

VOLUME /home/qmstr/config
VOLUME /var/qmstr/buildroot

ENTRYPOINT ["/usr/local/bin/fasten-reporter-standalone"]

