FROM golang:1.15-alpine
RUN mkdir /app
WORKDIR /app
COPY main assets templates ./
EXPOSE 8080
CMD [ "./main" ]
