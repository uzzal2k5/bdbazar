# Register/Login via auth-service
curl -X POST http://localhost:8081/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"seller@example.com", "password":"123456"}'

# Use returned JWT to create a product
curl -X POST http://localhost:8083/api/products \
  -H "Authorization: Bearer <jwt_token>" \
  -H "Content-Type: application/json" \
  -d '{"name":"T-Shirt","price":100,"description":"Cotton"}'
