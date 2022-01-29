
FROM python:slim

EXPOSE 8000


ENV PYTHONDONTWRITEBYTECODE=1


ENV PYTHONUNBUFFERED=1


COPY requirements.txt .
RUN python -m pip install -r requirements.txt

WORKDIR /app
COPY . /app


CMD ["gunicorn", "--bind", "0.0.0.0:8000", "docker.wsgi"]
