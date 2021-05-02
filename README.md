# Prototodo

To-Do List management cli with microservice & grpc

![image](https://user-images.githubusercontent.com/16697306/116808252-c4b8d100-ab72-11eb-927f-585102873cb3.png)
- Manager: Get user's request and send job to store & alarm
- Store: Get manager's job request and store to persistent storage
- Alarm: Get manager's job request and send alarm to user email when due date is reached

> All api should interact with grpc