# 1. Use a lightweight Python image as the base
FROM python:3.10-slim

# 2. Set the working directory inside the container
WORKDIR /app

# 3. Copy requirements and project files into the container
COPY requirements.txt /app/
COPY devops/ /app/

# 4. Install dependencies
#RUN python -m venv venv1 && source venv1/bin/activate && pip install --no-cache-dir -r requirements.txt
RUN python -m venv venv1 && /bin/bash -c "source venv1/bin/activate && pip install --upgrade pip --no-cache-dir -r requirements.txt"


# 5. Expose port 8000 for the application
EXPOSE 8000

# 6. Define the entry point for the container
ENTRYPOINT ["/bin/bash", "-c", "source venv1/bin/activate && python manage.py runserver 0.0.0.0:8000"]
