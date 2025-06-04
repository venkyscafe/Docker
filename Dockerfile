FROM ubuntu:latest

WORKDIR /my_app

COPY . /my_app

RUN apt-get update && apt-get install -y python3 python3-pip

CMD ["python3", "greetings.py"]
