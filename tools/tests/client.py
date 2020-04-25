import socket
import time

def add_one():
    data = b"<STX>Q,229,002.74,M,00,<ETX>16"
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect(("localhost", 8082))
    for _ in range(1):
        time.sleep(1)
        client.send(data)

def main():
    add_one()

if __name__ == "__main__":
    main()