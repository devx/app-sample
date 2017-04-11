FROM alpine:3.3
RUN apk -U add ca-certificates
ADD app-sample /usr/bin/app-sample
CMD app-sample
