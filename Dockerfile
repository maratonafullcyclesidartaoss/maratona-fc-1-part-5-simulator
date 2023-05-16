FROM golang:latest

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

USER nonroot

WORKDIR /app

COPY ./destinations .
COPY ./entity .
COPY ./queue .
COPY .env .
COPY .env.example .
COPY .gitignore .
COPY docker-compose.yaml .
COPY Dockerfile .
COPY go.mod .
COPY go.sum .
COPY README.md .
COPY simulator.go .
COPY sonar-project.properties .

RUN GOOS=linux go build simulator.go

CMD ["./simulator"]
