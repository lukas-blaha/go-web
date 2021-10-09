#!/usr/bin/env python3
# socket_echo_client.py
import socket
import sys

destination = '127.0.0.1'
base_port = 10000
connection_count = 8
chunk_length = 16 * 1024
sockets = []


def chunkstring(string, length):
    return (string[0+i:length+i] for i in range(0, len(string), length))

try:
  for i in range(connection_count):
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_address = (destination, base_port+i)
    #print('Connecting to {} port {}'.format(*server_address))
    sock.connect(server_address)
    sockets.append(sock)

  while True:
    #print("next")
    data = sys.stdin.buffer.read(connection_count * chunk_length)
    if len(data) == 0:
      break
    chunked = list(chunkstring(data, chunk_length))
    #print(str(chunked))
    for i in range(min(connection_count,len(chunked))):
      #print('sending {!r}'.format(chunked[i]))
      sockets[i].sendall(chunked[i])
      #print("sent {} bytes".format(len(chunked[i])))
#    del chunked[:connection_count]

finally:
  #print('Closing sockets')
  for i in range(connection_count):
    sockets[i].close()

# vim:autoindent:shiftwidth=2:tabstop=2:expandtab:
