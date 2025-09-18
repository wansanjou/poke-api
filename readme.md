# PokeAPI Proxy Service

RESTful API สำหรับดึงข้อมูล Pokémon จาก [PokeAPI](https://pokeapi.co/)  
รองรับ **User Register/Login** พร้อม **JWT Authentication** และ **Cache** เพื่อลดการเรียกซ้ำ

## 🛠 Tech Stack

- **Go 1.25 + Fiber**
- **MongoDB**
- **JWT Authentication**
- **In-memory Cache** (go-cache)
- **Docker + Docker Compose**

## 🚀 Quick Start

```bash
git clone https://github.com/<username>/poke-api.git
cd poke-api
สร้าง .env:
```

## ENV
```bash
PORT=8080
MONGO_URI=mongodb://mongo:27017
JWT_SECRET=aB3dE6FgH7jK8LmN9pQrStUvWxYz1234
JWT_SECRET ต้อง ≥ 32 ตัวอักษร
```

## Run Docker:
```bash
docker-compose up --build
Server: http://localhost:8080
MongoDB: mongodb://localhost:27017
``

## 📦 API Endpoints
Method
```bash
POST	/register
POST	/login	
GET	/api/pokemon/:name	
GET	/api/pokemon/:name/ability	
GET	/api/pokemon/random
```


## 📄 Examples
Register
http
POST /register
Content-Type: application/json

```bash
{
  "username": "ash",
  "password": "pikachu1234"
}
Response:

json
{
    "username": "ash",
    "created_at": "2025-09-18T11:42:15.846091927Z"
}
```

Login
http
POST /login
Content-Type: application/json

```bash
{
  "username": "ash",
  "password": "pikachu1234"
}
Response:

json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR..."
}
```

Get Pokémon
http
GET /api/pokemon/pikachu
Authorization: Bearer <token>
Response:

```bash
json
{
  "name": "pikachu",
  "types": ["electric"],
  "weight": 60,
  "abilities": ["static", "lightning-rod"]
}
```

Get Abilities
http
GET /api/pokemon/pikachu/ability
Authorization: Bearer <token>
Response:

```bash
json
{
  "abilities": ["static", "lightning-rod"]
}
```

## 📝 Notes
Cache Pokémon ข้อมูลละ 10 นาที เพื่อลดการเรียก PokeAPI ซ้ำ

JWT Token ต้องส่งใน Header ทุกครั้ง:
Authorization: Bearer <token>