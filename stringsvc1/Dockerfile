FROM golang:alpine AS build

RUN apk update
RUN apk add git
RUN apk add dos2unix

ENV GO111MODULE=on

WORKDIR /go/src/stringsvc
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/stringsvc main.go

FROM scratch
COPY --from=build /go/bin/stringsvc /go/bin/stringsvc

EXPOSE 8080

CMD ["/go/bin/stringsvc"]