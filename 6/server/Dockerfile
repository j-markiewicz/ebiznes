FROM golang:1.26 AS build

WORKDIR /server
COPY ./server/ /server/

RUN go build .

EXPOSE 8000
CMD [ "./projob5" ]
