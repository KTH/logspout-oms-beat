# logspout-oms-beat

A tiny helper docker image which outputs Bunyan formatted log messages, beats, 
at a given interval. It is used to generate information in order to monitor that
the log system is working properly. The image is tiny, slightly more than 2Mb.

I.e. the logs from this container is monitored from OMS with an alert in 
order to raise an alarm in case the log forwarding should fail.

## Configuration

There are a couple of optional variables with default values.

| Variable | Default | Description |
|----------|---------|-------------|
| LOGSPOUT_BEAT_LEVEL | 10 | Bunyan log level for output. |
| LOGSPOUT_BEAT_TIME | 60s | A Golang time Duration string for time between beats.|

## Running

Pre built images are available on docker hub as kthse/logspout-oms-beat.
They can be started as simply as: `docker run kthse/logspout-oms-beat:latest`.

In order to run as a supplement to logspout-oms, you most likely want to run
this as a global swarm service, like logspout-oms. Something like:

```
docker service create \
  --mode global \
  --restart-condition any \
  --restart-max-attempts 10 \
  --name="logspout-beat" \
  kthse/logspout-oms-beat:latest
```

## Development

This is a really tiny program and not much is expected to change. There is a 
build.sh script to help with compilation of the statically linked binary and
docker image. The options are passed to docker build. Hence,
`./build.sh -t mymage`
will build a docker image with tag myimage.

Run it with `docker run myimage`.
