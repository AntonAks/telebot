FROM python:3.8

# set environment variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1
ENV USE_DOCKER=true

WORKDIR /_parsers
COPY . .

RUN pip install --no-cache-dir -r requirements.txt


CMD ["python", "run.py"]