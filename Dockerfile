FROM golang:latest AS env

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
WORKDIR /app
RUN go mod download

FROM env AS build

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o ./training_sessions_receiver /app/cmd/training_sessions_receiver/main.go

FROM scratch
COPY --from=build /app/training_sessions_receiver /app/training_sessions_receiver

EXPOSE 4221
CMD ["/app/training_sessions_receiver"]