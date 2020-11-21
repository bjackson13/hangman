FROM golang:1.15-alpine
RUN mkdir /app
WORKDIR /app
COPY . ./
RUN go build main.go
EXPOSE 8080
CMD [ "./main" ]
