# Build

FROM golang:alpine AS build
ADD . /app
RUN cd /app && go build -o monitor-backend

# Deploy

FROM alpine
WORKDIR /monitor-backend
COPY --from=build /app /monitor-backend

EXPOSE 8081
ENTRYPOINT ["./monitor-backend"]
