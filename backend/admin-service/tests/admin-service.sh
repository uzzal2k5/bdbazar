curl -X POST http://localhost:8090/api/admins/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Admin User",
    "username": "admin123",
    "password": "StrongPassword123",
    "email": "admin@example.com",
    "mobile": "01712345678"
  }'


curl -X POST http://localhost:8090/api/admins/register \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJyb2xlIjoic3VwZXJhZG1pbiIsImV4cCI6MTc1NDIzMjcxNH0.1o_t7Y_mkIYM7yIn2_7znFb-3IFSKP2dSuxWUGAPcGc" \
  -d '{
    "name": "Admin User",
    "username": "admin123",
    "password": "StrongPassword123",
    "email": "admin@example.com",
    "mobile": "01712345678"
  }'


/api/admins/spadm/login
curl -X POST http://localhost:8090/api/admins/spadm/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "superadmin@example.com",
    "password": "StrongSuperPassword123"
  }'



SUPERUSER_NAME=Super Admin
SUPERUSER_USERNAME=superadmin
SUPERUSER_PASSWORD=StrongSuperPassword123
SUPERUSER_EMAIL=superadmin@example.com
SUPERUSER_MOBILE=01711111111

":"superadmin","username":"superadmin"},"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpX

{"superadmin":{"email":"superadmin@example.com","id":"1","roleVCJ9.eyJ1c2VyX2lkIjoxLCJyb2xlIjoic3VwZXJhZG1pbiIsImV4cCI6MTc1NDIyODYzNn0.yuU9qzKV0QyvHRbBfftNXU4WuiEUqa5isX8_xM0SI-E"}

admin-service uzzal$ curl -X POST http://localhost:8090/api/admins/register \
>   -H "Content-Type: application/json" \
>   -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJyb2xlIjoic3VwZXJhZG1pbiIsImV4cCI6MTc1NDIzMjcxNH0.1o_t7Y_mkIYM7yIn2_7znFb-3IFSKP2dSuxWUGAPcGc" \
>   -d '{
>     "name": "Admin User",
>     "username": "admin123",
>     "password": "StrongPassword123",
>     "email": "admin@example.com",
>     "mobile": "01712345678"
>   }'
{"error":"Invalid or expired token"}