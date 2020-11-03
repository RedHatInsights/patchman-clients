FROM centos:8

RUN yum module -y install go-toolset

ENV GOPATH=/go
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org
ENV PATH=/go/bin:$PATH

RUN mkdir -p /go/src/{inventory,rbac,vmaas}

ADD inventory/go.*  /go/src/inventory/
ADD rbac/go.*       /go/src/rbac/
ADD vmaas/go.*      /go/src/vmaas/

RUN cd /go/src/inventory && go mod download && \
    cd /go/src/vmaas && go mod download && \
    cd /go/src/vmaas && go mod download

ADD /inventory   /go/src/inventory
ADD /rbac        /go/src/rbac
ADD /vmaas       /go/src/vmaas

ADD run_tests.sh /go/src/

WORKDIR /go/src
