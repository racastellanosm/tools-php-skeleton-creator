#--- BUILDING STAGE ---#
FROM golang:1.22-alpine as builder
ARG SKAFFOLD_GO_GCFLAGS

WORKDIR /opt

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /opt/bin/equationlabs-php-cli ./equationlabs/cli.go

#--- PRODUCTION IMAGE ---#
FROM gcr.io/distroless/cc as production
COPY --from=builder /opt/bin/equationlabs-php-cli /equationlabs-php-cli

CMD ["/equationlabs-php-cli"]