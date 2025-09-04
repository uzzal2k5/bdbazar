🔹 1. Register auth-service at /api/auth

        curl -i -X POST http://localhost:8002/services \
          --data name=auth-service \
          --data url=http://auth-service:8080

        curl -i -X POST http://localhost:8002/services/auth-service/routes \
            --data paths[]=/api/auth

curl http://localhost:8060/api/auth/login

<!-- Test -->
#/api/auth/login - POST
curl -X POST http://localhost:8060/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "identifier": "uzzal@example.com",
    "password": "strongpassword123"
  }'


curl http://auth-service:8080/health

http://localhost:8080/


192.168.0.100



curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "identifier": "uzzal@example.com",
    "password": "strongpassword123"
  }'