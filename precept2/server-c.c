/*****************************************************************************
 * Name: 
 * NetId:
 *
 * Description: A simple TCP server using the C socket API.   Publishes
 *              an IPV4 IP address.
 *
 * Usage:       ./server-c [server port]
 * 
 * Example:     ./server-c 8999 
 * 
 *****************************************************************************/

#include <sys/socket.h>
#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#define QUEUE_LENGTH 10
#define RECV_BUFFER_SIZE 2048

/* 
 * Open socket and wait for client to connect.
 * Print received message to stdout.
 * Return 0 on success, non-zero on failure.
*/

int main(int argc, char **argv) {
  // get port number from command line
  char *server_port = argv[1];
  if (argc != 2) {
    fprintf(stderr, "Usage: ./server-c [port]\n");
    exit(EXIT_FAILURE);
  }

  // create a socket - IPV4 (AF_INET) and TCP (SOCK_STREAM)
  int socket_fd = 0;
  if ((socket_fd = socket(AF_INET, SOCK_STREAM, 0)) == -1) {
    perror("Error creating socket");
    exit(EXIT_FAILURE);
  }

  // configure the server address - use network byte ordering
  struct sockaddr_in server_address;   // from  <netinet/in.h>
  memset(&server_address, '0', sizeof(server_address));  // initialize serv_addr
  server_address.sin_family      = AF_INET;              // IPV4
  server_address.sin_addr.s_addr = htonl(INADDR_ANY);    // binds socket to all interfaces
                                                         // not just localhost   
  server_address.sin_port        = htons(atoi(server_port)); // network byte orders

  // bind socket to address (IP + port)
  if (bind(socket_fd, (struct sockaddr*)&server_address, sizeof(server_address)) == -1) {
    perror("Error setting socket options");
    exit(EXIT_FAILURE);
  }
  
  // listen on socket
  if (listen(socket_fd, QUEUE_LENGTH) == -1) {
    close(socket_fd);
    perror("Error listening on socket");
    exit(EXIT_FAILURE);
  }

  while (1) {
    // look for a client to connect
    int connection_fd = accept(socket_fd, (struct sockaddr*) NULL, NULL); 

    // receive a message from a client
    char message[RECV_BUFFER_SIZE];
    int bytes_received;
    if ((bytes_received = recv(connection_fd, message, RECV_BUFFER_SIZE, 0)) == -1) {
      close(connection_fd);
      perror("Error receiving from client");
      break;
    }

    // output the received data on stdout
    if (fwrite(message, sizeof(char), bytes_received, stdout) == -1 || fflush(stdout) == -1) {
      close(connection_fd);
      perror("Error writing to stdout");
      break;
    }

    // all done
    close(connection_fd);
  }
}
