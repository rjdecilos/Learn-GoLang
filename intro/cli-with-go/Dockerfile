FROM golang:1.16-alpine as dev
# my dev environment
WORKDIR /work

FROM golang:1.16-alpine as build
# build my go application
WORKDIR /videos
COPY ./videos/* /videos/
RUN go build -o videos

FROM alpine as runtime
# run buildfile in alpine
COPY --from=build /videos/videos /usr/local/bin/videos
COPY ./videos/videos.json /
COPY run.sh /
RUN chmod +x /run.sh
ENTRYPOINT [ "./run.sh" ]