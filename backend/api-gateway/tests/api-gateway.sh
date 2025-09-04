#auth
curl -X POST http://localhost:8000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "identifier": "uzzal@example.com",
    "password": "strongpassword123"
  }'
