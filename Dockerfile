FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main .
EXPOSE 3375
CMD ["/app/main"]
