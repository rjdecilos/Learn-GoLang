FROM golang:1.16-alpine as dev
# my dev environment
WORKDIR /work

FROM golang:1.16-alpine as build
# build my go application
WORKDIR /app
COPY ./app/* /app/
RUN go build -o app

FROM alpine as runtime
# run buildfile in alpine
COPY --from=build /app/app /
CMD ./app