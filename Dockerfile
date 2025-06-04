#From Base Image
FROM ubuntu:latest

#Create a default Working Directory
WORKDIR /my_app

#Copy the Python script into the app directory
COPY . /my_app

#Install Python and pip
RUN apt-get update && apt-get install -y python3 python3-pip

#Command to run the Python script
CMD ["python3", "greetings.py"]
