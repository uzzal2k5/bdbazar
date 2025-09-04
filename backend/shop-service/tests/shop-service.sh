

#API List
GET /api/shops
GET /api/shops/search
GET /api/shops/:id
GET /api/shops/dashboard

POST /api/shops/

PUT /api/shops/:id
DELETE /api/shops/:id


GET /api/shops/dashboard?owner_id=1

{
  "total_shops": 5,
  "approved_shops": 3,
  "blocked_shops": 1,
  "recent_shops_count": 2
}

✅ 1. Create a Shop

curl -X POST http://localhost:8084/api/shops \
  -H "Content-Type: application/json" \
  -d '{
    "name": "BDBazar Mart",
    "description": "A marketplace for local vendors.",
    "owner_id": 1
  }'

✅ 2. Update a Shop
curl -X PUT http://localhost:8084/api/shops/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "BDBazar Mart Updated",
    "description": "Updated description",
    "is_approved": true,
    "is_blocked": false
  }'

✅ 3. Get All Shops
curl http://localhost:8084/api/shops

✅ 4. Get Shop by ID
curl http://localhost:8084/api/shops/1

✅ 5. Delete a Shop
curl -X DELETE http://localhost:8084/api/shops/1


✅ 6. Search Shop by Name
curl "http://localhost:8084/api/shops/search?name=bazar"


TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InV6emFsQGV4YW1wbGUuY29tIiwiZXhwIjoxNzUzMDI4MzgyLCJpZCI6MywibW9iaWxlIjoiMDE5MTU4NDg0MTgiLCJyb2xlcyI6WyJzZWxsZXIiXX0.Lmnzugnltet_NCcbuEU67AMMHswiOi0yhhfKBZtCLSA"
curl -X POST http://localhost:8084/api/shops/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "name": "BDBazar Mart",
    "description": "A marketplace for local vendors."

  }'


echo $TOKEN | cut -d '.' -f2 | base64 -d

curl -X POST http://localhost:8080/api/auth/login   -H "Content-Type: application/json"   -d '{
    "identifier": "uzzal@example.com",
    "password": "strongpassword123"
  }'

sarah@example.com