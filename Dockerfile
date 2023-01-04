FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/DanielPickens/Scanly

ENV USER 
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET WcseCQOLZBDQ5Q7a

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://\Daniel@localhost:5432/Scanly?sslmode=disable

WORKDIR /go/src/github.com/DanielPickens/Scanly

RUN godep go build

EXPOSE 8888
CMD ./Scanly