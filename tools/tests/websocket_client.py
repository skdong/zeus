# pip install websocket-client

from websocket import create_connection
import json
ws = create_connection("ws://localhost:8080/api/ws/winds")
req = dict(devices=["Q"], limit=50, enable=True, close=False)
msg = json.dumps(req)
ws.send(msg)