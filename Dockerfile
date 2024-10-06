# syntax=docker/dockerfile:1
FROM golang:1.22 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --from=build-stage /entrypoint /entrypoint
COPY --from=build-stage /app/assets /assets
COPY --from=build-stage /app/css /css
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]