STEP 1: Create greeting.py python file.

STEP 2: Create Dockerfile with FROM, WORKDIR, COPY, RUN, CMD.

STEP 3: run the docker command to build the image from dockerfile -> "docker built -t v3nkys/file-to-image:latest ." (. -> represents the Dockerfile location, in this case its in working directory).

STEP 4: run the docker command to execute/run the image on a temporary container -> "docker run v3nkys/file-to-image:latest".

STEP 5: run the docker command to push the image to docker hub -> "docker push v3nkys/file-to-image:latest".
