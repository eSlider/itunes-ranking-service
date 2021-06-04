FROM golang:latest AS build
WORKDIR /app
COPY . .

# Need to compile SQlite
ENV CGO_ENABLED=1

RUN go get -d -v ./
RUN go build -a -o go-service .

#FROM scratch AS runtime
#COPY --from=build /app ./
EXPOSE 8080/tcp
ENTRYPOINT ["./go-service"]
