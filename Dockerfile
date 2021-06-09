FROM golang:latest AS build
WORKDIR /app
COPY . .

# Need to compile SQlite
ENV CGO_ENABLED=1

RUN go get -d -v ./
RUN go build -ldflags='-extldflags "-lm -lstdc++ -static  -w -s"' -tags sqlite_omit_load_extension -a -o go-service .

FROM scratch AS runtime
COPY --from=build /app ./
EXPOSE 8080/tcp
ENTRYPOINT ["./go-service"]
