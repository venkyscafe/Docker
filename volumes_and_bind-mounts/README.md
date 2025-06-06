Theory Part

Example 1: In a typical use case, an Nginx container logs user information such as usernames, IP addresses, and accessed URLs into a log file. These logs are useful for auditing and debugging. However, when the container is terminated or crashes, all files inside the container (including logs) are lost—unless persistent storage is used.
	
Example 2: Imagine a backend container generates data in formats like JSON, XML, or YAML on a daily basis. A frontend container is expected to read and display this data to the user. Without persistent storage, if the backend container goes down, all the generated records are lost, impacting data integrity and user experience.
	
Example 3: Suppose a cron job runs on the host machine and generates daily data files (JSON/XML/YAML/HTML). We may want an application running in a Docker container to access and display this data. This requires the container to access a shared directory from the host.
	
To address the above daily use cases, Docker introduced the concepts of volumes and bind mounts. While both serve the same purpose of providing persistent storage, volumes offer better security compared to bind mounts, as bind mounts require exposing specific host directories (e.g., /app) to the container.
	
		-> Bind Mounts: As the name suggests, a bind mount creates a direct link between a specific directory on the host machine and a specific directory inside the container. Any changes made within the container to that directory are reflected on the host, and vice versa. Even if the container stops or is removed, the data remains intact on the host. This makes bind mounts useful for scenarios where the same host directory needs to be reused across containers, ensuring data persistence.
		
		-> Volumes: A Docker volume is a storage mechanism managed by the Docker engine. It typically resides on the host machine (e.g., local disk or EC2 instance) and can be mounted to one or more containers, allowing data to persist across container restarts or deletions. Volumes are ideal for high-performance, read/write-intensive workloads and are easier to manage, back up, and migrate compared to bind mounts. They also offer better security and abstraction, as they're not directly linked to host file paths.

There is a subtle but important difference between the two ways to define mounts in Docker:

-v (short syntax): A concise way to define mounts, suitable for simple use cases.
	
 CMD -> docker run -v myvolume:/app/data mycontainer
	
--mount (long syntax): A more verbose and readable option, especially helpful for complex scenarios.

 CMD -> docker run --mount type=volume,source=myvolume,target=/app/data mycontainer

While both serve the same core purpose—attaching storage to containers—--mount is preferred in production environments for its clarity, explicitness, and flexibility.

Practical Part

	Step 1: Using the Docker command below, we can list out the volumes 
		
		CMD -> docker volume ls

	Step 2: Use the Docker command below to create a volume

		CMD -> docker volume create <VOLUME_NAME> 
		EX  -> docker volume create venky_volume

	Step 3: Using the Docker command below, we can get the specific <VOLUME_NAME> details

		CMD -> docker volume inspect <VOLUME_NAME>
		EX  -> docker volume inspect venky_volume
		
	Step 4: Using the below Docker command to remove/delete a specific or multiple <VOLUME_NAME(S)>

		CMD -> docker volume rm <VOLUME_NAME1> <VOLUME_NAME2>
		EX  -> docker volume rm venky_volume

	Step 5: Create a simple Docker file with Ubuntu as base image to create an Image and then a container to attach it to the volume <VOLUME_NAME>. 

		CMD -> FROM <IMAGE_NAME>:<TAG>
		EX  -> FROM ubuntu:latest
		
	Step 6: Now, build the image from the Docker file with the below command

		CMD -> docker build -t <IMAGE_NAME>:<TAG> .
		EX  -> docker build -t volume_image .
		
	Step 7: Run the below command to mount the volume with the --mount command to create more verbosity in -d detach mode

		CMD -> docker run -d --mount source=<VOLUME_NAME>,target=<APP_DIR> --name <CONTAINER_NAME> <IMAGE_NAME>:<TAG> (DEFAULT:READ-WRITE)
		EX  -> docker run -d --mount source=venky_volume,target=/app --name volume_container volume_image:latest
		
		CMD -> docker run -d --mount source=<VOLUME_NAME>,target=<APP_DIR>,<READONLY> --name <CONTAINER_NAME> <IMAGE_NAME>:<TAG> (READONLY MODE)
		EX  -> docker run -d --mount source=venky_volume,target=/app,readonly --name volume_container volume_image:latest
		
	Step 8: Using the inspect command, we can find out the container's volume details that we have created 

		CMD -> docker inspect <CONTAINER_NAME>
		EX  -> docker inspect volume_container

  	Step 9:	We can access the data created in the /app directory of the container from the host machine, but this path will require root privileges to access

   		CMD -> sudo -s

    and followed by

		CMD -> cd /var/lib/docker/volumes/<VOLUME_NAME>/_data
        	EX  -> cd /var/lib/docker/volumes/venky_volume/_data
	 
  	Step 10: We can see the same data from the container with the following command

    		CMD -> docker exec -it <CONTAINER_NAME> ls -ltr <APP_DIR>
      		EX  -> docker exec -it volume_container ls -ltr /app
		or
      		CMD -> docker exec <CONTAINER_NAME> ls -ltr <APP_DIR>
      		EX  -> docker exec volume_container ls -ltr /app

![image](https://github.com/user-attachments/assets/38c6ac32-dc89-4ac0-8f6d-a5e676fb80bb)

