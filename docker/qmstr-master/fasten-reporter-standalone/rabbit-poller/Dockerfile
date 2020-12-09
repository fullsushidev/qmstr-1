FROM python:3-alpine

ENV RABBIT_HOST="my-rabbit"
ENV RABBIT_PORT="5672"
ENV RABBIT_USERNAME="guest"
ENV RABBIT_PASSWORD="guest"

RUN pip install pika

COPY modules/reporters/fasten-reporter-standalone/rabbit-poller/rabbit_poller.py /rabbit_poller.py
RUN chmod +x /rabbit_poller.py

CMD ["sh", "-c", "python /rabbit_poller.py --host $RABBIT_HOST --port $RABBIT_PORT --username $RABBIT_USERNAME --password $RABBIT_PASSWORD"]
