#!/usr/bin/env python3
# socket_echo_server.py

from __future__ import print_function
import socket
import sys

def eprint(*args, **kwargs):
  print(*args, file=sys.stderr, **kwargs)

listen = '127.0.0.1'
base_port = 10000
connection_count = 8
chunk_length = 16 * 1024
sockets = []
recv_retries = 16

try:
  for i in range(connection_count):
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_address = (listen, base_port+i)
    eprint('starting up on {} port {}'.format(*server_address))
    sock.bind(server_address)
    sock.listen(1)
    sockets.append({ "server": sock})

  for i in range(connection_count):
    eprint("Waiting for connection {}".format(i))
    connection, client_address = sockets[i]["server"].accept()
    eprint("Accepted connection {} from {}".format(i, client_address))
    sockets[i]["connection"] = connection
    sockets[i]["client_address"] = client_address

  while True:
    for i in range(connection_count):
      data = sockets[i]["connection"].recv(chunk_length)
      #eprint('received on {} {!r}'.format(i, data))
      retry = recv_retries
      while chunk_length != len(data):
        eprint("Received {} bytes".format(len(data)))
        data2 = sockets[i]["connection"].recv(chunk_length-len(data))
        data += data2
        retry -= 1
        if retry == 0: break
      if retry != recv_retries: eprint("{} retries left".format(retry))
      if data: sys.stdout.buffer.write(data)
      else:
        eprint('no data from', sockets[i]["client_address"])
        sys.exit(1)
        break

finally:
  for i in range(connection_count):
    sockets[i]["connection"].close()

# vim:autoindent:shiftwidth=2:tabstop=2:expandtab:
