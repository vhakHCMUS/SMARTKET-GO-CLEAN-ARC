# SMARTKET GO CLEAN ARCHITECTURE - MVP

Backend API cho hệ thống SMARTKET được xây dựng với Go, Gin Framework và Clean Architecture.

## 🏗️ Kiến trúc

Dự án tuân theo **Clean Architecture** với các layer:

```
smartket-go-clean-arc/
├── domains/              # Domain layer (Entities & Interfaces)
├── infrastructure/       # Infrastructure layer (Database, External Services)
├── services/            # Business logic implementations
├── presentation/        # Presentation layer (HTTP Handlers)
├── api/                 # Routes & Middlewares
├── bootstrap/           # Dependency injection
├── lib/                 # Shared utilities
└── migration/           # Database migrations
```

## 🚀 Chức năng MVP

### 1. Auth Module ✅
- FR-Auth-01: Đăng ký qua Email
- FR-Auth-02: Đăng nhập Email + Password
- FR-Auth-06: Đăng xuất

### 2. Product/Search Module ✅
- FR-Search-01: Định vị thủ công (nhập địa chỉ text)
- FR-Search-02: Hiển thị danh sách sản phẩm
- FR-Search-03: Tìm kiếm theo tên
- FR-Search-04: Lọc cơ bản (giá, danh mục)

### 3. Purchase/Cart Module ✅
- FR-Purchase-01: Xem chi tiết sản phẩm
- FR-Purchase-05: Thêm vào giỏ
- FR-Purchase-06: Quản lý giỏ (nhóm theo shop)
- FR-Purchase-07: Đặt hàng đơn giản
- FR-Purchase-08: Thanh toán COD

### 4. Order Module ✅
- FR-Order-01: Xem danh sách đơn hàng
- FR-Order-04: Xác nhận nhận hàng (nhập mã đơn)

### 5. Profile Module ✅
- FR-Profile-02: Cập nhật thông tin cơ bản
- FR-Profile-05: Lịch sử đơn hàng

### 6. Merchant Module ✅
- FR-Merchant-01: Đăng ký merchant
- FR-Merchant-02: Đăng nhập merchant
- FR-Merchant-03: Upload sản phẩm
- FR-Merchant-04: Xem hàng tồn
- FR-Merchant-06: Xác nhận redeem
- FR-Merchant-10: Xem đơn hàng mới

## 📋 Yêu cầu

- Go 1.17+
- PostgreSQL 12+
- Docker (optional)

## 🔧 Cài đặt

### 1. Clone repository

```bash
git clone https://github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC.git
cd SMARTKET-GO-CLEAN-ARC
```

### 2. Cài đặt dependencies

```bash
go mod download
```

### 3. Cấu hình Database

Tạo file `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASS=your_db_password
DB_NAME=smartket
JWT_SECRET=your-secret-key
PORT=8080
```

### 4. Chạy migrations

```bash
# Sử dụng sql-migrate hoặc tool khác
sql-migrate up
```

### 5. Chạy server

```bash
go run server.go
```

Server sẽ chạy tại `http://localhost:8080`

## 📚 API Endpoints

### Auth APIs

```
POST   /api/auth/register        - Đăng ký user
POST   /api/auth/login           - Đăng nhập
POST   /api/auth/logout          - Đăng xuất (requires token)
GET    /api/auth/profile         - Xem profile (requires token)
PUT    /api/auth/profile         - Cập nhật profile (requires token)
```

### Merchant APIs

```
POST   /api/merchant/register    - Đăng ký merchant
POST   /api/merchant/login       - Đăng nhập merchant
GET    /api/merchant/profile     - Xem profile merchant (requires token)
PUT    /api/merchant/profile     - Cập nhật profile (requires token)
```

### Product APIs

```
GET    /api/products/search      - Tìm kiếm sản phẩm
GET    /api/products/:id         - Xem chi tiết sản phẩm

# Merchant only (requires merchant token)
POST   /api/merchant/products    - Tạo sản phẩm
GET    /api/merchant/products    - Xem danh sách sản phẩm của shop
PUT    /api/merchant/products/:id - Cập nhật sản phẩm
DELETE /api/merchant/products/:id - Xóa sản phẩm
```

### Cart APIs (requires token)

```
GET    /api/cart                 - Xem giỏ hàng
POST   /api/cart/add             - Thêm vào giỏ
PUT    /api/cart/items/:id       - Cập nhật số lượng
DELETE /api/cart/items/:id       - Xóa khỏi giỏ
POST   /api/cart/clear           - Xóa toàn bộ giỏ
```

### Order APIs (requires token)

```
# Customer
POST   /api/orders               - Tạo đơn hàng
GET    /api/orders               - Xem danh sách đơn hàng
GET    /api/orders/:id           - Xem chi tiết đơn hàng

# Merchant only
GET    /api/merchant/orders      - Xem đơn hàng của shop
POST   /api/merchant/orders/redeem - Xác nhận redeem đơn hàng
```

## 🔐 Authentication

API sử dụng JWT Bearer Token authentication.

**Request Header:**
```
Authorization: Bearer <your_jwt_token>
```

## 📝 Ví dụ Request/Response

### 1. Đăng ký User

**Request:**
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "name": "John Doe",
    "phone": "0123456789"
  }'
```

**Response:**
```json
{
  "data": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe",
    "phone": "0123456789",
    "role": "customer",
    "is_active": true,
    "created_at": "2024-11-11T10:00:00Z"
  }
}
```

### 2. Đăng nhập

**Request:**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

**Response:**
```json
{
  "data": {
    "access_token": "eyJhbGc...",
    "refresh_token": "eyJhbGc...",
    "user": {
      "id": 1,
      "email": "user@example.com",
      "name": "John Doe",
      "role": "customer"
    }
  }
}
```

### 3. Tìm kiếm sản phẩm

**Request:**
```bash
curl -X GET "http://localhost:8080/api/products/search?keyword=bread&category=bakery&max_price=50000"
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "merchant_id": 1,
      "name": "Bánh mì baguette",
      "description": "Bánh mì tươi ngon",
      "category": "bakery",
      "orig_price": 20000,
      "sale_price": 15000,
      "discount": 25,
      "stock": 50,
      "images": "https://example.com/bread.jpg",
      "expiry_date": "2024-11-12T00:00:00Z",
      "is_active": true
    }
  ],
  "total": 1
}
```

### 4. Thêm vào giỏ hàng

**Request:**
```bash
curl -X POST http://localhost:8080/api/cart/add \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "merchant_id": 1,
    "quantity": 2
  }'
```

### 5. Tạo đơn hàng

**Request:**
```bash
curl -X POST http://localhost:8080/api/orders \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "merchant_id": 1,
    "delivery_address": "123 Nguyen Hue, Q1, TPHCM",
    "payment_method": "COD",
    "notes": "Giao trước 5pm",
    "items": [
      {
        "product_id": 1,
        "quantity": 2
      }
    ]
  }'
```

**Response:**
```json
{
  "data": {
    "id": 1,
    "order_code": "ORD17314560001234",
    "total_amount": 30000,
    "status": "pending",
    "payment_method": "COD",
    "payment_status": "unpaid",
    "items": [
      {
        "product_id": 1,
        "quantity": 2,
        "price": 15000,
        "subtotal": 30000,
        "product_name": "Bánh mì baguette"
      }
    ]
  }
}
```

### 6. Merchant xác nhận redeem

**Request:**
```bash
curl -X POST http://localhost:8080/api/merchant/orders/redeem \
  -H "Authorization: Bearer <merchant_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "order_code": "ORD17314560001234"
  }'
```

## 🐳 Docker

```bash
# Build
docker-compose build

# Run
docker-compose up

# Stop
docker-compose down
```

## 🧪 Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## 📦 Database Schema

### Users Table
- id, email, password, name, phone, role, is_active, created_at, updated_at

### Merchants Table
- id, user_id, shop_name, shop_address, phone, latitude, longitude, description, is_verified, is_active

### Products Table
- id, merchant_id, name, description, category, orig_price, sale_price, discount, stock, images, expiry_date, is_active

### Orders Table
- id, user_id, merchant_id, order_code, total_amount, status, payment_method, payment_status, delivery_address, pickup_time, completed_at, notes

### Order Items Table
- id, order_id, product_id, quantity, price, subtotal, product_name

### Carts & Cart Items Tables
- Quản lý giỏ hàng của user

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License.

## 👥 Authors

- VhakHCMUS - [GitHub](https://github.com/vhakHCMUS)

## 🙏 Acknowledgments

- Original template from [dipeshdulal/clean-gin](https://github.com/dipeshdulal/clean-gin)
