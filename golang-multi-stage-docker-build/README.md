Theory Part

	We can make Docker images even better by using something called Distroless images. Using Distroless images, we can improve the efficiency of multi-stage Docker builds. These are very lightweight because they don’t include basic tools like a shell or package managers—they only contain what’s needed to run the app. This makes the image: Smaller in size, More secure & Faster to run. 

	For example, a language like Go (Golang) works well with Distroless because Go creates a self-contained binary. That means it doesn’t need a separate runtime like Python or Java. So you can run the app directly without extra files or dependencies. When we combine multi-stage builds with Distroless images, we get the best of both worlds—clean, minimal, and highly secure Docker images that are ready for production.

	Normally, Docker is considered lightweight, but the images we build can become quite heavy. This is because we often use base images like Ubuntu and install extra tools like Python, Java, or pip, even though we only need the runtime environment to run our code. As a result, our images end up containing a lot of unnecessary files and dependencies, which increases their size and makes them less efficient.

	To solve this problem, Docker introduced a concept called multi-stage builds. With this approach, we can separate the build process from the final image. For example, we can use one stage to build the frontend, another to build the backend, and another to prepare the database. Then, in the final stage, we only copy the required files from each of these stages as aliases into a clean and minimal image that includes only the things needed to run the app, like Python or Java runtime.

	This helps us keep the final image small, secure, and optimised for production.

Practical Part

	Step 1: Go to the folder golang-multi-stage-docker-build/dockerfile-without-multistage to find the Dockerfile without multistage implementation, where you will see that the size of the image is very heavy when built with the command.
	
		-> docker build -t image-without-multistage .

	Step 2: Run the Docker command to create a container

		-> docker run -d --name <CONTAINER_NAME> <IMAGE_ID>

	Step 3: Now enter into the container with the below docker command, followed by the golang command to run the calculator.go

		-> docker exec -it <CONTAINER_NAME> bash
		or
		-> docker exec -it <CONTAINER_NAME> /bin/bash
	
		-> go run calculator.go
	
	Step 4: NOTE DOWN THE IMAGE SIZE OF THE image-without-multistage

	Step 5: Now go to the folder golang-multi-stage-docker-build to find the Dockerfile with multistage implementation, where you will see the lightweight size of the image when built with the command.

		-> docker build -t image-with-multistage .
	
	Step 6: Run the Docker command to create the container

		-> docker run -d --name <CONTAINER_NAME> <IMAGE_ID>
	
	Step 7: Now, enter into the container with the Docker command below, followed by the Golang command to run the calculator.go

		-> docker exec -it <CONTAINER_NAME> bash
		or
		-> docker exec -it <CONTAINER_NAME> /bin/bash
	
		-> go run calculator.go
	
	Step 8: NOTE DOWN THE IMAGE SIZE OF THE image-with-multistage

	Conclusion: Upon comparing the image size of an image with multistage against an image without multistage, we can notice a huge variance of 32,275%.

![image](https://github.com/user-attachments/assets/af6c5648-b79d-41ee-a8e9-8aa537abd7ec)

