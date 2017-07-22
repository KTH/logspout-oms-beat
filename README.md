# logspout-oms-beat

A tiny helper docker image which outputs Bunyan formatted log messages, beats, 
at a given interval. It is used to generate information in order to monitor that
the log system is working properly.

I.e. the logs from this container is monitored from OMS with an alert in 
order to raise an alarm in case the log forwarding should fail.

## Configuration

There are a couple of optional variables with default values.

|----------|---------|-------------|
| Variable | Default | Description |
|----------|---------|-------------|
| LOGSPOUT_BEAT_LEVEL | 10 | Bunyan log level for output. |
| LOGSPOUT_BEAT_TIME | 60s | A Golang time Duration string for time between beats.|

## Development

This is a really tiny program and not much is expected to change. There is a 
build.sh script to help with compilation of the statically linked binary and
docker image. The options are passed to docker build. Hence,
`./build.sh -t mymage`
will build a docker image with tag myimage.

Run it with `docker run myimage`.
