FROM ubuntu:latest

RUN apt-get -y update && apt-get -y install python3 && apt-get -y install pip
RUN apt update
RUN apt-get -y install git
RUN git clone https://github.com/LeATeP/my_whatever

RUN pip install psycopg2-binary && pip install configparser

ENV PATH="/usr/bin:/bin:/game/hub"
ENV POSTGRES_HOST=postgres
ENV POSTGRES_USER=worker
ENV POSTGRES_PASSWORD=123
ENV POSTGRES_DB=sql

WORKDIR /my_whatever/game/hub

