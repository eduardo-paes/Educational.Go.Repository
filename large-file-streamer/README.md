# File Transfer Server

This README provides an overview of the code for a simple file transfer server and client written in Go. The server listens on a specified port for incoming file transfers, while the client generates and sends a random file to the server.

## Table of Contents

- [File Transfer Server](#file-transfer-server)
  - [Table of Contents](#table-of-contents)
  - [1. Introduction](#1-introduction)
  - [2. Prerequisites](#2-prerequisites)
  - [3. Code Structure](#3-code-structure)
  - [4. Usage](#4-usage)

## 1. Introduction

This Go code is designed to demonstrate a basic file transfer system. It consists of a server and a client.

- The **server** listens on a specified port (default is `:8080`) and awaits incoming file transfers using the TCP protocol. It uses a simple protocol to receive files, specifying the size of the file and reading the file in chunks.

- The **client** is a simple program that generates a random file of a specified size and sends it to the server.

## 2. Prerequisites

Before running this code, make sure you have the following prerequisites:

- Go installed on your system.
- A basic understanding of Go programming.

## 3. Code Structure

The code is organized as follows:

- `main.go`: The main entry point of the application.
- `FileServer` struct: Responsible for setting up the server and handling incoming file transfers.
- `start()`: Starts the server and listens for incoming connections.
- `readLoop(conn net.Conn)`: Reads files sent by clients in chunks.
- `sendFile(size int) error`: Generates a random file and sends it to the server.
- The `main()` function starts both the server and the client. The client generates and sends a file to the server after a 1-second delay.

## 4. Usage

To use this code:

1. Clone or download the repository to your local machine.

2. Navigate to the directory where the code is located.

3. Run the server:

   ```sh
   go run main.go
   ```

   The server will start listening for incoming file transfers on port `:8080` by default.

4. After the server is running, it will automatically accept incoming file transfers.

5. To send a file to the server, a client is already configured to send a file of 123,456 bytes after a 1-second delay. You can customize the file size in the `sendFile` function.

6. Observe the server's output to see the progress of file transfers.

7. The server will print out the received file data and the number of bytes received.

8. The client will print out the number of bytes written.

9. You can customize the transfer parameters by modifying the constants defined in the code, such as `PORT`, `TRANSFER_RATE`, and `TRANSFER_TYPE`.

**Note**: This code is intended for educational purposes and may require additional error handling and security measures for production use.
