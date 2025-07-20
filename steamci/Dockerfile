FROM golang:bookworm

COPY * .
ADD templates templates

RUN go build -o /bin/steamci .
RUN ls

FROM debian

COPY --from=0 /bin/steamci /bin/steamci

WORKDIR /home/steam
RUN apt update
RUN apt install -y lib32gcc-s1 curl tar

RUN curl -sqL "https://steamcdn-a.akamaihd.net/client/installer/steamcmd_linux.tar.gz" | tar zxvf -
RUN chmod +x steamcmd.sh
RUN /home/steam/steamcmd.sh +quit

RUN rm -rf /var/lib/apt/lists/*

CMD ["/bin/steamci"]