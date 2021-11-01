FROM registry.access.redhat.com/ubi8

RUN dnf module -y install go-toolset

ENV GOPATH=/go
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org
ENV PATH=/go/bin:$PATH

ADD /rbac        /go/src/rbac
RUN cd /go/src/rbac && go mod download

ADD /inventory   /go/src/inventory
RUN cd /go/src/inventory && go mod download

ADD /vmaas       /go/src/vmaas
RUN cd /go/src/vmaas && go mod download

ADD run_tests.sh /go/src/

WORKDIR /go/src
