FROM golang@sha256:9dd2625a1ff2859b8d8b01d8f7822c0f528942fe56cfe7a1e7c38d3b8d72d679

RUN apk --no-cache --quiet add tini-static

ENTRYPOINT ["/sbin/tini-static", "--"]
CMD ["/sbin/ticker"]