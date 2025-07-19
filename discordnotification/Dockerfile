FROM golang:alpine

COPY * .
ADD templates templates

RUN go build -o /bin/discordnotification .
RUN ls

FROM alpine
COPY --from=0 /bin/discordnotification /bin/discordnotification
CMD ["/bin/discordnotification"]