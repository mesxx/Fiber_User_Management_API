FROM golang:1.22.0

RUN addgroup app && adduser --system --group app

USER app

WORKDIR /app

COPY go.* .

USER root

RUN mkdir -p /nonexistent/.cache/go-build && \
    chown -R app:app /nonexistent/.cache && \
    chown -R app:app .

USER app

RUN go mod download

COPY . .

EXPOSE 5000

CMD go run main.go