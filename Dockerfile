FROM ubuntu

RUN apt update
RUN apt-get install ca-certificates -y

WORKDIR /app

CMD ["./test"]