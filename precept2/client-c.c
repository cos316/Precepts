/*****************************************************************************
 * Name: 
 * NetId:
 *
 * Description: A simple client using C socket API.  Sends up to 
 *              SEND_BUFFER_SIZE bytes or EOF.   Limited to IP address
 *              of server (e.g., 10.0.2.15  or 127.0.0.1) instead of domain 
 *              name (e.g., mycomputer.cs.princeton.edu or localhost)
 * 
 * Usage:       Usage: ./client-c [server IP] [server port] < [message]
 *              (Assumes server is running.)
 *
 * Example:     ./client-c 127.0.0.1 8999 < somefile.txt
 * 
 *****************************************************************************/

#include <sys/socket.h>
#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <netdb.h>
#include <unistd.h>
#include <arpa/inet.h>

#define SEND_BUFFER_SIZE 2048

/* Open socket and send message from stdin.
 * Return 0 on success, non-zero on failure
*/

int main(int argc, char **argv) {
  // get IP and port number from command line
  char *server_ip   = argv[1];
  char *server_port = argv[2];
  if (argc != 3) {
    fprintf(stderr, "Usage: ./client-c [server IP] [server port] < [message]\n");
    exit(EXIT_FAILURE);
  }

  // create a socket - IPV4 (AF_INET) and TCP (SOCK_STREAM)
  int socket_fd = 0;
  if ((socket_fd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
    perror("Error creating socket");
    exit(EXIT_FAILURE);
  }

  // configure the server address - use network byte ordering
  struct sockaddr_in server_address;   // from  <netinet/in.h>
  memset(&server_address, '0', sizeof(server_address));  // initialize serv_addr
  server_address.sin_family      = AF_INET;              // IPV4
  server_address.sin_port        = htons(atoi(server_port)); // network byte orders
  if (inet_pton(AF_INET, server_ip, &server_address.sin_addr) <= 0) {
    close(socket_fd);
    fprintf(stderr, "Unable to parse IP address: %s\n", server_ip);
    exit(EXIT_FAILURE);
  } 

  // connect to server
  if (connect(socket_fd, (struct sockaddr *)&server_address, sizeof(server_address)) != 0) {
    close(socket_fd);
    perror("Error connecting to socket");
    exit(EXIT_FAILURE);
  } 

  // read message from stdin - up to SEND_BUFFER_SIZE bytes or EOF
  char message[SEND_BUFFER_SIZE];
  int bytes_read, bytes_sent;
  bytes_read = fread(message, sizeof(char), SEND_BUFFER_SIZE, stdin);

  // check for error
  if (ferror(stdin)) {
    close(socket_fd);
    perror("Error reading from stdin");
    return 1;
  }

  // call send  - uo to SEND_BUFFER_SIZE bytes
  int bytes_to_send;
  if (bytes_read > SEND_BUFFER_SIZE) 
    bytes_read = SEND_BUFFER_SIZE;
  if ((bytes_sent = send(socket_fd, message, bytes_read, 0)) == -1) {
    close(socket_fd);
    perror("Error sending message");
    return 1;
  }

  close(socket_fd);
  return 0;
}
