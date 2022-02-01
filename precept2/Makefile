# Simple Makefile for compiling the client and server examples.

all: client-c server-c

client-c: client-c.c
	gcc client-c.c -o client-c

server-c: server-c.c
	gcc server-c.c -o server-c

clean:
	rm -f client-c server-c *.o
