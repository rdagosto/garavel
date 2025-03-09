FROM golang:1.24

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.43.0 && \
    wget -qO- https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/ && \
    chmod +x /usr/local/bin/migrate

ENV PATH="/go/bin:/usr/local/bin:$PATH"

CMD ["air"]
