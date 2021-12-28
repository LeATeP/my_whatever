FROM ubuntu:latest

RUN apt-get update && apt-get -y install python3
RUN apt-get update && apt-get -y install pip && pip install psycopg2-binary
RUN apt-get update && \
    apt-get -qy full-upgrade && \
    apt-get install -qy curl && \
    curl -sSL https://get.docker.com/ | sh
RUN apt-get -y install vim

# RUN git clone https://github.com/LeATeP/my_whatever

ENV GAME=/project/game_docker
ENV PATH=$GAME:$GAME/hub:/project/docker/script:$PATH

ENV POSTGRES_HOST=postgres
ENV POSTGRES_USER=postgres
ENV POSTGRES_DB=sql
ENV POSTGRES_PASSWORD=123

WORKDIR /project/game_docker