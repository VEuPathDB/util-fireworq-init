FROM alpine:3

ARG SETUP_UTIL_URL=https://github.com/VEuPathDB/util-fireworq-init/releases/download/v1.0.0/queue-setup.v1.0.0.x64.tar.gz
ARG FIREWORQ_URL=https://github.com/fireworq/fireworq/releases/download/v1.4.1/fireworq_linux_amd64.zip

# Install utils, fetch fireworq, fetch queue setup
RUN apk add --no-cache wget tar zip \
    && wget ${FIREWORQ_URL} -O fireworq.zip \
    && wget ${SETUP_UTIL_URL} -O setup.tar.gz \
    && unzip fireworq.zip && mv fireworq /usr/local/bin/fireworq && rm fireworq.zip \
    && tar -xf setup.tar.gz && mv setup /usr/local/bin/setup && rm setup.tar.gz

COPY queues.yml queues.yml

CMD ["sh", "-c", "fireworq & lpid=$! sleep 5s; setup; wait $lpid"]
