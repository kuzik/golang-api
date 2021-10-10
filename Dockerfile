## We'll choose the incredibly lightweight
## Go alpine image to work with
FROM golang:latest AS builder

## We create an /app directory in which
## we'll put all of our project code
RUN mkdir /build
ADD . /build
WORKDIR /build
## We want to build our application's binary executable
RUN GOOS=linux CGO_ENABLED=0 go build -o server

## the lightweight scratch image we'll
## run our application within
FROM alpine:latest AS production
## We have to copy the output from our
## builder stage to our production stage
COPY --from=builder /build .
## we can then kick off our newly compiled
## binary exectuable!!
CMD ["./server"]