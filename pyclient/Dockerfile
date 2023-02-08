FROM python:3

RUN mkdir /tmp/client
WORKDIR /tmp/client

ADD requirements.txt ./
RUN pip install -r requirements.txt
ADD client.py ./

CMD ["python3", "client.py"]
