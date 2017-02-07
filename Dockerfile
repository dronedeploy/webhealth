FROM golang:1.7
MAINTAINER DroneDeploy <admin@dronedeploy.com>
LABEL REPO="https://github.com/dronedeploy/webhealth"


COPY main.go /go/src/github.com/dronedeploy/webhealth/main.go
COPY cmd     /go/src/github.com/dronedeploy/webhealth/cmd/

RUN go get github.com/spf13/cobra
RUN go install -a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo github.com/dronedeploy/webhealth
ENTRYPOINT webhealth

# if it got this far, label it with the git hash
ARG GIT_HASH
LABEL GIT_HASH=$GIT_HASH
ENV GIT_HASH=$GIT_HASH
