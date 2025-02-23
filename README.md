# cert-generator

curl -X GET http://localhost:9393/communities

curl -X POST http://localhost:9393/communities -H "Content-Type: application/json" -d '{
    "name": "AWG UG Triangulo Mineiro",
    "organizer": "Rogerío Fontes"
}'

curl -X GET http://localhost:9393/events
curl -X GET http://localhost:9393/events -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5AZXhhbXBsZS5jb20iLCJyb2xlIjoidXNlciIsImV4cCI6MTc0MDMxODIwMn0.Z4Lp_H8Pb1CQdMtrdO8SORujOvjwjExNVc5rYLoHJTU"
curl -X POST http://localhost:9393/events -H "Content-Type: application/json" -d '{
    "name": "16 Meetup da AWG UG Triangulo Mineiro",
    "description": "Evento de segurança cibernética",
    "date": "2025-05-20",
    "local": "Online",
    "total_hours": 12,
    "community_id": 1
}'

curl -X POST http://localhost:9393/participants -H "Content-Type: application/json" -d '{
    "name": "João Silva",
    "email": "joao@email.com",
    "status": "pendente",
    "certificate": "",
    "event_id": 1
}'

curl -X POST http://localhost:9393/eventos/1/participantes -H "Content-Type: application/json" -d '{
    "nome": "Maria Souza",
    "email": "maria@email.com",
    "status": "pendente",
    "certificado": "",
    "evento_id": 1
}'

curl -X POST http://localhost:9393/eventos/1/participantes -H "Content-Type: application/json" -d '{
    "nome": "Rogério Fontes 1",
    "email": "rogerio@email.com",
    "status": "pendente",
    "certificado": "",
    "evento_id": 1
}'

curl -X GET http://localhost:9393/eventos/1/participantes

curl -X POST http://localhost:9393/certificados/evento/1

curl -X POST http://localhost:8080/certificados/evento/999

curl -X POST http://localhost:9393/certificates/event/1

curl -X GET http://localhost:9393/certificate-participants/validate?code=e3a81743-9b68-4670-9eb4-2b035b8a1fac

curl -X POST http://localhost:9393/users/register -H "Content-Type: application/json" -d '{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "123456"
}'

curl -X POST http://localhost:9393/users/login -H "Content-Type: application/json" -d '{
  "email": "john@example.com",
  "password": "123456"
}'

//TODO - criar crud de comunidade -ok 
//TODO- colocar swagger 
//Colocar JWT
//Criar uma validação de QRCODE -ok
//Arrumar UTF-8
go get github.com/skip2/go-qrcode


====
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

📌 Explicação:

swag init → Comando usado para gerar a documentação.
gin-swagger → Middleware para exibir o Swagger UI no browser.
swagger/files → Arquivos necessários para o Swagger UI funcionar.

go mod tidy
swag init -g cmd/main.go --output ./docs

rm -rf docs

go run cmd/main.go
http://localhost:9393/swagger/index.html

swag init -g cmd/main.go -o docs

go get -u github.com/dgrijalva/jwt-go
openssl rand -base64 32
go run generate_secret.go

curl -X POST http://localhost:9393/login -H "Content-Type: application/json" -d '{
  "username": "admin",
  "password": "123456"
}'

curl -X GET http://localhost:9393/events -H "Authorization: Bearer SEU_TOKEN_AQUI"
‣癥挭牥⵴敧敮慲潴⵲灡੩