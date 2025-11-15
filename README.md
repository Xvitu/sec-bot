# Security Bot
Um chatbot simples para assuntos de cyber sec

# Funcionalidades

# Como Rodar Localmente
- Instale o docker e docker compose
- Rode o comando
- 
`` 
docker compose build && docker compose up -d  
``

``
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name chat-updates-queue
aws --endpoint-url=http://localhost:4566 sqs receive-message --queue-url http://localhost:4566/000000000000/chat-updates-queue
``

``
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main processor.go
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main receiver.go
``

``
sam local invoke ChatConsumerFunction --event event.json --docker-network sam-local
sam local start-api
``