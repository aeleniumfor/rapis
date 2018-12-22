FROM alpine:latest as build
ADD https://github.com/sosedoff/pgweb/releases/download/v0.9.12/pgweb_linux_amd64.zip  ./
RUN apk add unzip && unzip pgweb_linux_amd64.zip && mv pgweb_linux_amd64 pgweb

FROM alpine:latest
COPY --from=build ./pgweb ./pgweb
EXPOSE 8080
CMD [ "./pgweb", "--sessions" ,"--bind=0.0.0.0", "--listen=8080" ]
