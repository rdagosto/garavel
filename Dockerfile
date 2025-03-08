FROM golang:1.21

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.43.0

ENV PATH="/go/bin:$PATH"

CMD ["air"]
