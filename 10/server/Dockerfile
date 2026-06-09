FROM golang:1.26

WORKDIR /server
COPY . /server/

RUN go build .

EXPOSE 8000
CMD [ "./ebiz10" ]
