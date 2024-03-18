FROM python:alpine3.11

WORKDIR /app

COPY . .

RUN apk update

RUN apk add git

RUN pip install -r requirements.txt

CMD ["python","-u", "./cmd/model/main.py"]