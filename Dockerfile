FROM golang:latest as build
ENV GO111MODULE=on
ENV PATH /go/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "/go/src/app" "/go/bin" && chmod -R 777 "/go"
WORKDIR /go/src/app
COPY . .
RUN go mod init screenshot-site; go mod tidy
RUN go build

FROM chromedp/headless-shell:latest
RUN apt-get update; apt install dumb-init -y
ENTRYPOINT ["dumb-init", "--"]
COPY --from=build /go/src/app/screenshot-site/ /tmp
WORKDIR /tmp
CMD ./screenshot-site
