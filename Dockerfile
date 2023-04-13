FROM golang:alpine AS build-env

# Set up dependencies
ENV PACKAGES git build-base file

# Set working directory for the build
WORKDIR /go/src/github.com/vknowable/pocket-core

# Install dependencies
RUN apk add --update $PACKAGES
RUN apk add linux-headers

# Add source files
COPY . .

# Stable version (RC-0.9.2)
RUN git checkout 785379d

# Build pocket-core
RUN go build -o /go/bin/pocket /go/src/github.com/vknowable/pocket-core/app/cmd/pocket_core/main.go

# ------ Final image ------ #
FROM alpine:3.17.0

# Install ca-certificates
RUN apk add --update ca-certificates jq curl
WORKDIR /

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/pocket /usr/bin/pocket

# Check if build successful
CMD ["pocket"]
