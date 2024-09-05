FROM golang:1.23.0

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN git config --global --add safe.directory '*' # Fix for vcs issue
RUN go mod tidy
