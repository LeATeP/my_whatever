FROM ubuntu:latest

# RUN git clone https://github.com/LeATeP/my_whatever

ENV GAME=/project/game_docker
ENV PATH=$GAME:$GAME/hub:/project/docker/script:$PATH

ENV POSTGRES_HOST=postgres
ENV POSTGRES_USER=postgres
ENV POSTGRES_DB=sql
ENV POSTGRES_PASSWORD=123

WORKDIR /project/game_docker