FROM ubuntu

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update -y && apt install -y vim git curl wget nodejs npm golang && mkdir -p /app

RUN go get -u github.com/gopherjs/gopherjs
RUN go get golang.org/dl/go1.12.16
RUN /root/go/bin/go1.12.16 download

COPY entrypoint.sh /

WORKDIR /app
COPY package.json .
RUN npm install


COPY Gruntfile.js .
COPY .babelrc .
COPY src src

RUN ["/entrypoint.sh"]
