FROM golang:1.14.3 AS build
LABEL maintainer="Gary Ellis <gary.luis.ellis@gmail.com>"

ARG VERSION

ENV NAME cloud-infrastructure-sdk
ENV WORKDIR_BASE /github.com/garyellis

WORKDIR $WORKDIR_BASE/$NAME

COPY . $WORKDIR_BASE/$NAME
RUN package=github.com/garyellis/$NAME/pkg/cli VERSION=$VERSION && \
    BUILD_DATE="-X '${package}.BuildDate=$(date)'" && \
    GIT_COMMIT="-X ${package}.GitCommit=$(git rev-list -1 HEAD)" && \
    _VERSION="-X ${package}.Version=$VERSION" && \
    FLAGS="$GIT_COMMIT $_VERSION $BUILD_DATE" && \
    export GOOS=linux GOARCH=amd64 && go build -o /release/${NAME}-${VERSION}_${GOOS}-${GOARCH} -ldflags "${FLAGS}" && \
    export GOOS=darwin GOARCH=amd64 && go build -o /release/${NAME}-${VERSION}_${GOOS}-${GOARCH} -ldflags "${FLAGS}"

FROM ubuntu
COPY --from=build /release/ /release/
