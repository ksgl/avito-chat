FROM ubuntu:18.04

ENV PGSQLVER 10
ENV DEBIAN_FRONTEND 'noninteractive'

RUN echo 'Europe/Moscow' > '/etc/timezone'

RUN apt-get -y update
RUN apt install -y gcc git wget
RUN apt install -y postgresql-$PGSQLVER

RUN wget https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz
RUN tar -xvf go1.11.2.linux-amd64.tar.gz
RUN mv go /usr/local

ENV GOROOT /usr/local/go
ENV GOPATH /opt/go
ENV PATH $GOROOT/bin:$GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR $GOPATH/avito-chat
COPY . .

EXPOSE 9000

USER postgres

RUN /etc/init.d/postgresql start &&\
    psql --echo-all --command "CREATE USER ksu WITH SUPERUSER PASSWORD 'pswd';" &&\
    createdb -O ksu avito_chat_db &&\
    psql --dbname=avito_chat_db --echo-all --command 'CREATE EXTENSION IF NOT EXISTS citext;' &&\
    psql avito_chat_db -f $GOPATH/avito-chat/database/scheme.sql &&\
    /etc/init.d/postgresql stop

RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGSQLVER/main/pg_hba.conf
RUN echo "listen_addresses='*'" >> /etc/postgresql/$PGSQLVER/main/postgresql.conf
RUN echo "synchronous_commit = off" >> /etc/postgresql/$PGSQLVER/main/postgresql.conf
RUN echo "shared_buffers = 512MB" >> /etc/postgresql/$PGSQLVER/main/postgresql.conf

EXPOSE 5432

USER root

RUN go build $GOPATH/avito-chat/cmd/main.go
CMD service postgresql start && ./main