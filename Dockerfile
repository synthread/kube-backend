FROM alpine:3.19

RUN apk add --no-cache tini
ENTRYPOINT ["/sbin/tini", "--"]

RUN mkdir -p /app

ADD out/kube-backend /app/kube-backend
ADD static/ /app/static

CMD ["/app/kube-backend"]
