import json
import requests




jsonData = {
    "id": 0,
    "params": ["hello world"],
    "method": "HelloService.Hello"
}
response = requests.post("http://localhost:1234/jsonrpc", json=jsonData)
print(response.text)