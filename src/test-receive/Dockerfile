# Build

FROM golang:alpine AS build
ADD . /app
RUN cd /app && go build -o test-receive

# Deploy

FROM alpine
WORKDIR /test-receive
COPY --from=build /app /test-receive
ENTRYPOINT ["./test-receive"]
