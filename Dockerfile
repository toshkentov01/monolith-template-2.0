# workspace (GOPATH) configured at /go
FROM golang:1.18.2 as builder


#
RUN mkdir -p $GOPATH/src/github.com/sardortoshkentov/monolith-template-2.0
WORKDIR $GOPATH/src/github.com/sardortoshkentov/monolith-template-2.0

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/monolith_template /


FROM alpine
COPY --from=builder monolith_template .
RUN mkdir config
COPY ./config/rbac_model.conf ./config/rbac_model.conf
COPY ./config/models.csv ./config/models.csv
# COPY ./templates/ ./templates/
ENTRYPOINT ["/monolith_template"]