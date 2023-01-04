FROM golang:1.19 as build
ARG BUILD_REF

COPY . /service

WORKDIR /service
RUN go build -ldflags "-X main.build=${BUILD_REF}"

FROM alpine:latest
ARG BUILD_DATE
ARG BUILD_REF
COPY --from=build /service /service
WORKDIR /service
CMD ["./make"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
  org.opencontainers.image.title="make" \
  org.opencontainers.image.revision="${BUILD_REF}"
