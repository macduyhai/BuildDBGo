FROM golang:1.14 as builder

WORKDIR /app_adscreen

COPY . /app_adscreen

RUN go mod download

RUN GOOS=linux

RUN go build -o main

FROM ubuntu:16.04

WORKDIR /app_adscreen

COPY --from=builder /app_adscreen/main .

EXPOSE 80
ENV SET_CONTAINER_TIMEZONE false
ENV CONTAINER_TIMEZONE Asia/Ho_Chi_Minh
CMD ["/app_adscreen/main"]