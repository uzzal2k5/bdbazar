# Start PostgreSQL
docker-compose up -d db # Make sure your .env contains the correct DATABASE_URL.

# Run Auth Service
go run cmd/main.go


#/api/auth/register - POST
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "johndoe@example.com",
    "mobile": "01712345678",
    "password": "strongpassword123",
    "roles": ["buyer", "seller"]
  }'

curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "admin",
    "email": "admin@bdbazar.com",
    "mobile": "501715519132",
    "password": "strongpassword123",
    "roles": ["admin"]
  }'


  curl -X POST http://localhost:8080/api/auth/register \
    -H "Content-Type: application/json" \
    -d '{
      "name": "uzzal",
      "email": "uzzal@example.com",
      "mobile": "01915848418",
      "password": "strongpassword123",
      "roles": ["buyer"]
    }'


#/api/auth/login - POST
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "identifier": "uzzal@example.com",
    "password": "strongpassword123"
  }'



#/api/auth/refresh - POST
curl -X POST http://localhost:8080/api/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "-G8xwG-Gfuqh9VS0rIF5eTW9_bpjSSLWD3XFUrMnagI="
  }'


#/api/auth/logout - POST
curl -X POST http://localhost:8080/api/auth/logout \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "rrtW-QkRiQkeStgUNerBJIWQR-9poFY8SLSfqgJ47U4="
  }'

#/api/user/profile
curl -i -X GET http://localhost:8080/api/user/profile \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InV6emFsQGV4YW1wbGUuY29tIiwiZXhwIjoxNzUyMzk2MTczLCJpZCI6MywibW9iaWxlIjoiMDE5MTU4NDg0MTgiLCJyb2xlcyI6WyJidXllciJdfQ.2bqKpdtmvlmJjcywMDYvjpNn19Wf_hkJfIhqBxgWkGc"

#/api/seller/dashboard
curl -X GET http://localhost:8080/api/seller/dashboard \
  -H "Authorization: Bearer <your_jwt_access_token>"


