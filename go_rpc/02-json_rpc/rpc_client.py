import json
import socket




# json rpc client
request = {
    "id": 0,
    "params": ["hello world"],
    "method": "HelloService.Hello"
}


client = socket.create_connection(("localhost",1234))
client.sendall(json.dumps(request).encode())

# 获取服务器返回的数据
rsp = client.recv(1024)
rsp = json.loads(rsp.decode())
print(rsp)

