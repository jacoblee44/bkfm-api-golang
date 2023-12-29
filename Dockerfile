FROM golang:alpine AS build

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o app .

FROM alpine

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

WORKDIR /app

COPY --from=build /tmp/app/app /app/

RUN chmod +x /app/app

EXPOSE 8080

ENTRYPOINT ["/app/app"]