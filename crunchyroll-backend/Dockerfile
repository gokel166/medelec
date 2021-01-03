FROM python:3.8.5-alpine
ENV PYTHONUNBUFFERED REP_PASS
WORKDIR /app
COPY requirements.txt /app/requirements.txt
RUN pip install -r requirements.txt
COPY . /app

CMD python manage.py runserver 0.0.0.0:5000
