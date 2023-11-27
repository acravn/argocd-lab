FROM golang:1.20.0-alpine
LABEL authors="andrew ravn"

CMD mkdir /app

WORKDIR /app

COPY main.go .

EXPOSE 3333
CMD ["go", "run", "main.go"]