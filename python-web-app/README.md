To work with this Python web application, it's important to have a basic understanding of how a Django web application is structured—even if it's just the skeleton.

After installing Python, you can install Django using pip:

	pip install django

Once installed, the django-admin tool becomes available. It's similar in concept to ansible-galaxy (used in Ansible for generating role structures). While ansible-galaxy scaffolds the folder and file structure for roles, django-admin helps you generate the initial project layout, including settings, manage scripts, and application directories.

For example:

	django-admin startproject devops

This command creates a project with essential files like:

	devops/
	├── manage.py
	└── devops/
		├── __init__.py
		├── settings.py
		├── urls.py
		├── asgi.py
		└── wsgi.py
		
Later, you can create Django apps within the project using:

	python manage.py startapp demo

This will create a structure like:

	demo/
	├── admin.py
	├── apps.py
	├── models.py
	├── tests.py
	├── views.py
	└── migrations/

Main things to remember are

	settings.py is crucial for managing environments, security, integrations, and behavior of your Django app. For a DevOps engineer, understanding how to adjust or externalize these settings (via environment variables, config maps, or secrets) is key when deploying Django apps in cloud or containerized environments.
	
		-> python-web-app/devops/devops/settings.py
		
	urls.py file defines URL patterns that map specific URLs to corresponding views (i.e., functions or classes that handle the request and return a response) and ensures that every incoming HTTP request reaches the right view, making it a core component of Django’s MTV architecture (Model-Template-View).
	
		-> python-web-app/devops/devops/urls.py
 
	views.py contains the python code which is executed against the rendering the templates such as a demo_site.html file. views.py file in a Django application is where the core business logic of your web app resides. It works as the bridge between the models (data) and the templates (UI).
 
		-> python-web-app/devops/demo/views.py
 
	demo_site.html will be served on the browser using URL
	
		-> python-web-app/devops/demo/templates/demo_site.html

Step 1: Create Dockerfile with FROM, WORKDIR, COPY, RUN, EXPOSE, ENTRYPOINT
	
	-> FROM : using lightweight python image instead of ubuntu
	-> RUN : we are directly installing dependencies directly since its a python Image
		-> creating a Python virtual environment (venv1)
		-> Activating the virtual environment
		-> Installing the dependencies listed in requirements.txt inside the virtual environment.
	-> ENTRYPOINT : 
		-> Changes the default shell to Bash for running commands.
		-> Activates the virtual environment.
		-> Starts the Django development server (manage.py runserver), binding to all network interfaces (0.0.0.0).
		
Step 2: run the docker command to build the image from dockerfile (./dot represents the Dockerfile location, in this case its in current directory).

	-> docker build .
	
Step 3: run the docker command to create the container or to run the image with a temperory container
	-> docker run -it -p 8000:8000 --name <CONTAINER_NAME> <IMAGE_ID>
								or
	-> docker run -d -p 8000:8000 --name <CONTAINER_NAME> <IMAGE_ID>
	-> docker start <CONTAINER_NAME>
								or
	-> docker run -it -p 8000:8000 <IMAGE_ID>
	
Step 4: access the webpage using <EC2_INSTAINCE_IP_ADDRESS>:<PORT_NUMBER>/demo/
