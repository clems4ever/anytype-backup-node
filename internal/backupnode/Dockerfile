# any-sync-coordinator
FROM golang:1.20-bullseye AS any-sync-coordinator
LABEL MAINTAINER="Clément Michaud <clement.michaud34@gmail.com>"

WORKDIR /usr/app

RUN git clone --depth 1 --branch v0.2.6 https://github.com/anyproto/any-sync-coordinator.git
RUN cd any-sync-coordinator && make deps && make build

CMD ["/usr/app/any-sync-coordinator/bin/any-sync-coordinator", "-c", "/etc/anytype/coordinator.yml"]

# any-sync-node
FROM golang:1.20-bullseye AS any-sync-node
LABEL MAINTAINER="Clément Michaud <clement.michaud34@gmail.com>"

WORKDIR /usr/app

RUN git clone --depth 1 --branch v0.2.8 https://github.com/anyproto/any-sync-node.git
RUN cd any-sync-node && make deps && make build
RUN mkdir /usr/app/db

CMD ["/usr/app/any-sync-node/bin/any-sync-node", "-c", "/etc/anytype/sync_1.yml"]

# any-sync-filenode
FROM golang:1.20-bullseye AS any-sync-filenode
LABEL MAINTAINER="Clément Michaud <clement.michaud34@gmail.com>"

WORKDIR /usr/app

# We point to the fork until this PR is merged: https://github.com/anyproto/any-sync-filenode/pull/22
RUN git clone --depth 1 --branch v0.3.4-s3-creds https://github.com/clems4ever/any-sync-filenode.git
RUN cd any-sync-filenode && make deps && make build

CMD ["/usr/app/any-sync-filenode/bin/any-sync-filenode", "-c", "/etc/anytype/file_1.yml"]
