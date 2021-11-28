FROM ubuntu:latest

RUN apt-get update && \
    apt-get -y install python3 && \
    apt-get -y install pip

RUN pip install psycopg2-binary && \
    pip install configparser

RUN apt-get update && \
    apt-get -qy full-upgrade && \
    apt-get install -qy curl && \
    curl -sSL https://get.docker.com/ | sh

# RUN apt-get update && apt-get -y install git
# RUN git clone https://github.com/LeATeP/my_whatever

ENV PATH="/usr/bin:/bin:/my_whatever/game_docker:/my_whatever/game_docker/hub:my_whatever/docker/script"
ENV POSTGRES_HOST=postgres
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=123
ENV POSTGRES_DB=sql

WORKDIR /game

