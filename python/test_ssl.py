import ssl
import socket
import time
import argparse
import requests

URL = "https://benchmark-test-462.s3.us-west-2.amazonaws.com/persons.json"
PORT = 443
HOSTNAME = "benchmark-test-462.s3.us-west-2.amazonaws.com"

def measure_ssl_handshake_and_network_io_time(iterations):
    total_ssl_handshake_time = 0
    total_network_io_time = 0

    for _ in range(iterations):
        # Perform a GET request to the specified URL
        start_time_network_io = time.time()
        response = requests.get(URL)
        received_data = response.content
        network_io_time = (time.time() - start_time_network_io) * 1000
        total_network_io_time += network_io_time

        # Create a TCP socket and connect to the server
        client_socket = socket.create_connection((HOSTNAME, PORT))

        # Wrap the socket in an SSL context for the SSL handshake
        ssl_context = ssl.create_default_context()
        ssl_socket = ssl_context.wrap_socket(client_socket, server_hostname=HOSTNAME)

        # Measure the time taken for the SSL handshake
        start_time_ssl_handshake = time.time()
        ssl_socket.do_handshake()
        ssl_handshake_time = (time.time() - start_time_ssl_handshake) * 1000
        total_ssl_handshake_time += ssl_handshake_time

        # Close the connection
        ssl_socket.close()

    return (
        total_ssl_handshake_time / iterations,
        total_network_io_time / iterations,
    )

def main():
    parser = argparse.ArgumentParser(description="Measure SSL handshake and network I/O performance.")
    parser.add_argument(
        "-i",
        "--num_iterations",
        type=int,
        default=10,
        help="Number of iterations for testing.",
    )
    args = parser.parse_args()

    # Measure SSL handshake and network I/O time
    (
        average_ssl_handshake_time,
        average_network_io_time,
    ) = measure_ssl_handshake_and_network_io_time(args.num_iterations)

    print(f"Avg SSL Handshake Time ({args.num_iterations} iterations): {average_ssl_handshake_time:.6f} milliseconds")
    print(f"Avg Network I/O Time ({args.num_iterations} iterations): {average_network_io_time:.6f} milliseconds")

if __name__ == "__main__":
    # Target server information
    main()
