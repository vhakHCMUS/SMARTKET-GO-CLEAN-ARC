-- Seed data for development and testing
SET NAMES utf8mb4;

-- Insert sample merchant user
INSERT INTO users (email, password, name, phone, role) VALUES
('merchant@smartket.com', '$2a$10$abcdefghijklmnopqrstuvwxyz', 'Gia Lạc Minimart', '0901234567', 'merchant')
ON DUPLICATE KEY UPDATE email=email;

-- Insert merchant profile
INSERT INTO merchants (user_id, business_name, address, latitude, longitude, phone, is_verified)
SELECT id, 'Gia Lạc Minimart', 'Quận 1, TP.HCM', 10.7769, 106.7009, '0901234567', 1
FROM users WHERE email = 'merchant@smartket.com'
ON DUPLICATE KEY UPDATE business_name=business_name;

-- Insert sample products
INSERT INTO products (merchant_id, name, description, category, orig_price, sale_price, discount, stock, expiry_date, is_active) 
SELECT 
    m.id,
    'Mì Hảo Hảo Tôm Chua Cay',
    'Gói mì ăn liền hương vị tôm chua cay',
    'Thực phẩm & Đồ ăn',
    8000,
    6000,
    25,
    50,
    DATE_ADD(NOW(), INTERVAL 7 DAY),
    1
FROM merchants m WHERE m.business_name = 'Gia Lạc Minimart'
ON DUPLICATE KEY UPDATE name=name;

INSERT INTO products (merchant_id, name, description, category, orig_price, sale_price, discount, stock, expiry_date, is_active)
SELECT 
    m.id,
    'Cơm Bento Trứng Cuộn',
    'Hộp cơm bento với trứng cuộn Nhật Bản',
    'Thực phẩm & Đồ ăn',
    50000,
    35000,
    30,
    20,
    DATE_ADD(NOW(), INTERVAL 1 DAY),
    1
FROM merchants m WHERE m.business_name = 'Gia Lạc Minimart'
ON DUPLICATE KEY UPDATE name=name;

INSERT INTO products (merchant_id, name, description, category, orig_price, sale_price, discount, stock, expiry_date, is_active)
SELECT 
    m.id,
    'Bánh Mì Việt Nam',
    'Bánh mì thịt nguội truyền thống',
    'Bánh ngọt / Bánh mì',
    25000,
    18000,
    28,
    15,
    DATE_ADD(NOW(), INTERVAL 1 DAY),
    1
FROM merchants m WHERE m.business_name = 'Gia Lạc Minimart'
ON DUPLICATE KEY UPDATE name=name;

INSERT INTO products (merchant_id, name, description, category, orig_price, sale_price, discount, stock, expiry_date, is_active)
SELECT 
    m.id,
    'Sữa Tươi Vinamilk',
    'Hộp sữa tươi không đường 1L',
    'Sữa & sản phẩm từ sữa',
    32000,
    28000,
    13,
    30,
    DATE_ADD(NOW(), INTERVAL 5 DAY),
    1
FROM merchants m WHERE m.business_name = 'Gia Lạc Minimart'
ON DUPLICATE KEY UPDATE name=name;

INSERT INTO products (merchant_id, name, description, category, orig_price, sale_price, discount, stock, expiry_date, is_active)
SELECT 
    m.id,
    'Cà Phê Đen Đá',
    'Ly cà phê đen đá truyền thống',
    'Đồ uống',
    20000,
    15000,
    25,
    25,
    DATE_ADD(NOW(), INTERVAL 1 DAY),
    1
FROM merchants m WHERE m.business_name = 'Gia Lạc Minimart'
ON DUPLICATE KEY UPDATE name=name;
