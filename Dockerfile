FROM golang:1.21.3 AS build-stage
LABEL authors="lipe"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /planning-poker

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM golang:1.21.3-alpine AS build-release-stage
WORKDIR /

COPY --from=build-stage /planning-poker /planning-poker
COPY ./templates ./templates

EXPOSE 8080
ENV GIN_MODE=release

ENTRYPOINT ["/planning-poker"]