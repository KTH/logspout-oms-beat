FROM scratch

# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o beat .

ENV LOGSPOUT_BEAT_LEVEL=10
ENV LOGSPOUT_BEAT_TIME=60s

ADD beat /beat
CMD ["/beat"]
