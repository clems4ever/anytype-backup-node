# any-sync-coordinator
FROM golang:1.19.11-bullseye AS any-sync-coordinator
LABEL MAINTAINER="Clément Michaud <clement.michaud34@gmail.com>"

WORKDIR /usr/app

RUN git clone https://github.com/anyproto/any-sync-coordinator.git && cd any-sync-coordinator && git checkout v0.2.8
RUN cd any-sync-coordinator && make deps && make build

CMD ["/usr/app/any-sync-coordinator/bin/any-sync-coordinator", "-c", "/etc/anytype/coordinator.yml"]

# any-sync-node
FROM golang:1.19.11-bullseye AS any-sync-node
LABEL MAINTAINER="Clément Michaud <clement.michaud34@gmail.com>"

WORKDIR /usr/app

RUN git clone https://github.com/anyproto/any-sync-node.git && cd any-sync-node && git checkout v0.2.8
RUN cd any-sync-node && make deps && make build
RUN mkdir /usr/app/db

CMD ["/usr/app/any-sync-node/bin/any-sync-node", "-c", "/etc/anytype/sync_1.yml"]

# any-sync-filenode
FROM golang:1.19.11-bullseye AS any-sync-filenode
LABEL MAINTAINER="Clément Michaud <clement.michaud34@gmail.com>"

WORKDIR /usr/app

RUN git clone https://github.com/anyproto/any-sync-filenode.git && cd any-sync-filenode && git checkout v0.3.4
RUN cd any-sync-filenode && make deps && make build

CMD ["/usr/app/any-sync-filenode/bin/any-sync-filenode", "-c", "/etc/anytype/file_1.yml"]
