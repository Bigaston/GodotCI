FROM golang:alpine

COPY * .

RUN go build -o /bin/bunnyci .
RUN ls

FROM alpine

COPY --from=0 /bin/bunnyci /bin/bunnyci

CMD ["/bin/bunnyci"]