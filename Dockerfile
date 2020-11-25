FROM ubuntu

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update -y && apt install -y \
    curl wget \
    vim git \
    gnupg \
    dos2unix

RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list

RUN apt update -y && apt install -y \
    golang \
    nodejs \
    yarn 

RUN go get -u github.com/gopherjs/gopherjs
RUN go get golang.org/dl/go1.12.16
RUN /root/go/bin/go1.12.16 download

RUN mkdir -p /app
WORKDIR /app
COPY package.json .
RUN yarn install
RUN apt install -y 

COPY Gruntfile.js .
COPY .babelrc .
COPY src src

COPY entrypoint.sh /
ENTRYPOINT ["/entrypoint.sh"]

RUN dos2unix /app/* /app/*/* /entrypoint.sh