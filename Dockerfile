FROM golang:1.16-alpine AS build

WORKDIR /src
ADD . /src
RUN go mod download
RUN go build -o app ./cmd/main.go


FROM alpine

WORKDIR /app
COPY --from=build /src/app /app/
CMD ["./app"]
