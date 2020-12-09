#!/usr/bin/env python3
# Check connection to the RabbitMQ server
# Source: https://blog.sleeplessbeastie.eu/2017/07/10/how-to-check-connection-to-the-rabbitmq-message-broker/

import argparse
import time
import pika

# define and parse command-line options
parser = argparse.ArgumentParser(description='Check connection to RabbitMQ server')
parser.add_argument('--host', required=True, help='Define RabbitMQ server hostname')
parser.add_argument('--virtual_host', default='/', help='Define virtual host')
parser.add_argument('--port', type=int, default=5672, help='Define port (default: %(default)s)')
parser.add_argument('--username', default='guest', help='Define username (default: %(default)s)')
parser.add_argument('--password', default='guest', help='Define password (default: %(default)s)')
args = vars(parser.parse_args())

print(args)

# set amqp credentials
credentials = pika.PlainCredentials(args['username'], args['password'])
# set amqp connection parameters
parameters = pika.ConnectionParameters(host=args['host'], port=args['port'], virtual_host=args['virtual_host'], credentials=credentials)

# try to establish connection and check its status
while True:
    try:
        connection = pika.BlockingConnection(parameters)
        if connection.is_open:
            print('OK')
            connection.close()
            exit(0)
    except Exception as error:
        print('No connection yet:', error.__class__.__name__)
        time.sleep(5)
