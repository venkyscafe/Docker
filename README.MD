STEP 1: Create greeting.py python file.

STEP 2: Create Dockerfile with FROM, WORKDIR, COPY, RUN, CMD.

STEP 3: run the docker command to build the image from dockerfile (./dot represents the Dockerfile location, in this case its in working directory).

	-> docker built -t v3nkys/file-to-image:latest .

STEP 4: run the docker command to execute/run the image on a temporary container.

	-> docker run v3nkys/file-to-image:latest

STEP 5: run the docker command to login and then enter the password when it prompts.

	-> docker login -u username

STEP 6: run the docker command to push the image to docker hub.

	 -> docker push v3nkys/file-to-image:latest

STEP 7: perform these steps to push the changes to github and don't forget that entering username/password is deprecated. Hence creating tokens is adviced.

	-> Settings -> Developer Settings -> Personal access tokens -> Tokens (classic) -> Generate new token (classic) -> Set the NOTE, EXPERIRATION DATE, SELECT SCOPES (avoid selcting admin scopes) -> Generate token (Copy the token and save somewhere, as you won’t be able to see it again!)

	-> git init
	-> git add .
	-> git commit -m "msg"
	-> git branch -M main
	-> git remote add origin https://github.com/USER_NAME/REPO_NAME.git
	-> git push https://<your-token>@github.com/USER_NAME/REPO_NAME.git

STEP 8: edit .bashrc in your home dir with below changes, to see which branch you are in

	# Function to show current Git branch
	parse_git_branch() {
	  git rev-parse --abbrev-ref HEAD 2>/dev/null
	}

	if [ "$color_prompt" = yes ]; then
	    PS1='${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\w\[\033[33m\] $(parse_git_branch)\[\033[00m\]\$ '
	else
	    PS1='${debian_chroot:+($debian_chroot)}\u@\h:\w $(parse_git_branch)\$ '
	fi

STEP 9: change your global configs to push the code to github under your username otherwise it will be ubuntu

	o git config --global user.name "username"
	o git config --global user.email "email_address"

