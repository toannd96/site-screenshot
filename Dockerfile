FROM golang:latest as build
ENV GO111MODULE=on
ENV PATH /go/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "/go/src/app" "/go/bin" && chmod -R 777 "/go"

WORKDIR /go/src/app
COPY . /go/src/app

#RUN go build -o bin/screenshot-site
RUN go mod download

FROM chromedp/headless-shell:latest
RUN apt-get update; apt install dumb-init -y
ENTRYPOINT ["dumb-init", "--"]
COPY --from=build /go/src/app /app
WORKDIR /app
CMD /app/bin/screenshot-site
