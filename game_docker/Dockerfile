FROM ubuntu:latest

RUN apt-get -y update && apt-get -y install python3 && apt-get -y install pip
RUN pip install psycopg2-binary && pip install configparser

WORKDIR /game

