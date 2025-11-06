ARG IMAGE="go:latest"
FROM cgr.dev/jimmy.ungerman/${IMAGE}

WORKDIR /app
COPY . .

RUN go build -o server main.go

EXPOSE 8080
ENTRYPOINT ["./server"]