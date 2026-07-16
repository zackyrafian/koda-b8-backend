-- DROP TABLE IF EXISTS
--     product_previews,
--     product_images,
--     product_variants,
--     products,
--     categories,
--     brands, 
--     users, 
--     user_profiles, 
--     user_address, 
--     user_orders, 
--     user_carts, 
--     user_wishlists;

CREATE TABLE brands( 
  id BIGSERIAL PRIMARY KEY, 
  name VARCHAR (50) NOT NULL UNIQUE, 
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE categories( 
  id BIGSERIAL PRIMARY KEY, 
  name VARCHAR(50) NOT NULL UNIQUE, 
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE products ( 
  id BIGSERIAL PRIMARY KEY, 
  name VARCHAR(255) NOT NULL,
  brand_id BIGINT NOT NULL REFERENCES brands(id),
  category_id BIGINT NOT NULL REFERENCES categories(id),
  
  price BIGINT NOT NULL DEFAULT 0,
  discount INT NOT NULL DEFAULT 0,
  rating NUMERIC(2, 1) DEFAULT 0,
  stock INT NOT NULL DEFAULT 0,
  
  sold_out BIGINT NOT NULL DEFAULT 0, 
  description TEXT NOT NULL, 
  created_at TIMESTAMP DEFAULT NOW(), 
  updated_at TIMESTAMP 
);

CREATE TABLE product_variants( 
  id BIGSERIAL PRIMARY KEY, 
  name VARCHAR(40) NOT NULL, 
  product_id BIGINT REFERENCES products(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE product_images( 
  id BIGSERIAL PRIMARY KEY, 
  product_id BIGINT REFERENCES products(id) ON DELETE CASCADE,
  url TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE product_previews( 
  id BIGSERIAL PRIMARY KEY,
  product_id BIGINT REFERENCES products(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE users( 
  id BIGSERIAL PRIMARY KEY, 
  email VARCHAR(50) NOT NULL UNIQUE, 
  password VARCHAR(80) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_profiles( 
  id BIGSERIAL PRIMARY KEY, 
  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
  first_name VARCHAR(30) NOT NULL, 
  last_name VARCHAR(30),
  image_profile TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_address(
  id BIGSERIAL PRIMARY KEY, 
  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_carts(
  id BIGSERIAL PRIMARY KEY, 
  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE, 
  product_id BIGINT REFERENCES products(id) ON DELETE CASCADE, 
  quantity INT 
);

CREATE TABLE user_orders( 
  id BIGSERIAL PRIMARY KEY, 
  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
  product_id BIGINT REFERENCES products(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_wishlists(
  id BIGSERIAL PRIMARY KEY, 
  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
  product_id BIGINT REFERENCES products(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO brands(name)
VALUES 
('SoundWare'),
('LogiTech');

INSERT INTO categories(name)
VALUES 
('Audio'),
('Aksesoris Komputer');

INSERT INTO products(
    name,
    brand_id,
    category_id,
    price,
    stock,
    discount,
    rating,
    sold_out,
    description
)
VALUES
(
    'Headphone Wireless Premium',
    1,
    1,
    650000,
    120,
    0,
    4.8,
    156,
    'Nikmati pengalaman mendengar yang superior dengan Headphone Wireless Premium dari SoundWare. Dilengkapi dengan teknologi pemutusan kebisingan aktif (ANC) dan driver 40mm yang menghasilkan suara jernih serta bass yang dalam. Koneksi Bluetooth 5.2 memastikan koneksi stabil, sementara baterai berdaya tahan hingga 30 jam memungkinkan Anda menikmati musik sepanjang hari tanpa khawatir kehabisan daya. Desain ergonomis dengan bantalan empuk memberikan kenyamanan ekstra saat digunakan dalam jangka waktu lama.'
),
(
    'Mouse Gaming RGB Pro',
    2,
    2,
    350000,
    120,
    10,
    4.7,
    89,
    'Tingkatkan permainan Anda dengan Mouse Gaming RGB Pro dari LogiTech. Dilengkapi dengan sensor optik 16.000 DPI yang dapat disesuaikan, memberikan presisi dan kecepatan tinggi untuk berbagai genre game. Lampu RGB yang dapat dikustomisasi menambah estetika gaming setup Anda. Desain ergonomis dengan grip samping memastikan kenyamanan selama sesi permainan yang panjang. Mouse ini juga dilengkapi dengan 7 tombol yang dapat diprogram, memungkinkan Anda untuk menyesuaikan kontrol sesuai kebutuhan.'
);

INSERT INTO product_variants(product_id, name)
VALUES
(1, 'Hitam'),
(1, 'Putih'),
(1, 'Biru'),
(2, 'Hitam'),
(2, 'Putih');


INSERT INTO product_images(product_id, url)
VALUES
(1, '/headphone1.png'),
(1, '/headphone2.png'),
(1, '/headphone3.png');

SELECT
    p.id,
    p.name,
    b.name AS brand,
    c.name AS category,
    string_agg(pv.name, ', ' ORDER BY p.id) AS variant,
    p.price,
    p.description
FROM products p
JOIN brands b
    ON p.brand_id = b.id
JOIN categories c
    ON p.category_id = c.id
JOIN product_variants pv
    ON pv.product_id = p.id
GROUP BY 
    p.id,
    p.name,
    b.name,
    c.name,
    p.price;


