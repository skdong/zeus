import socket
import time

def add_one():
    data = [0x02, 0x51, 0x2C, 0x2C,
		0x30, 0x30, 0x30, 0x2E, 0x30,
		0x32, 0x2C, 0x4D, 0x2C, 0x30, 0x30, 0x2C,
		0x03, 0x32, 0x43, 0x0D, 0x0A]
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect(("localhost", 5003))
    for _ in range(600):
        time.sleep(1)
        client.send(bytes(data))

def main():
    add_one()

if __name__ == "__main__":
    main()