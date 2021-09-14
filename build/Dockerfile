FROM amazonlinux:2

RUN yum install -y golang-1.15.8-1.amzn2.0.1.x86_64 pkgconfig libaio-devel jq

ENV GOSUMDB=off
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOPATH=/go
ENV PATH=${PATH}:${GOPATH}/bin

ENV REPOSITORY=github.com/m-machida1108/prime-factorization
ARG WORKDIR=${GOPATH}/src/${REPOSITORY}

WORKDIR ${WORKDIR}

ENTRYPOINT [ "tail", "-f", "/dev/null" ]