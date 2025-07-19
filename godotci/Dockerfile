FROM ubuntu

ARG VERSION=4.3-stable

# Install required packages
RUN apt update
RUN apt install -y zip
RUN apt install -y curl
RUN apt install -y wget
RUN apt install -y fontconfig
RUN apt install -y git

RUN mkdir /godot

COPY download_godot.sh .
RUN sh download_godot.sh ${VERSION}
RUN rm download_godot.sh

# Clean APT
RUN rm -rf /var/lib/apt/lists/*

RUN godot --version