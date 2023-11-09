# Build

FROM golang:alpine AS build
ADD . /app
RUN cd /app && go build -o data-collector

# Deploy

FROM alpine
WORKDIR /data-collector
COPY --from=build /app /data-collector
ENTRYPOINT ["./data-collector"]
