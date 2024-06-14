CREATE TABLE category (
    id UUID PRIMARY KEY,
    slug VARCHAR UNIQUE NOT NULL, -- 'generated' as note
    name_uz VARCHAR(20) DEFAULT '',
    name_ru VARCHAR(20) DEFAULT '',
    name_en VARCHAR(20) DEFAULT '',
    active BOOLEAN DEFAULT TRUE,
    order_no INTEGER DEFAULT 0,
    parent_id UUID REFERENCES category(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE product (
    id UUID PRIMARY KEY,
    slug VARCHAR UNIQUE NOT NULL, -- 'generated' as note
    name_uz VARCHAR(20) DEFAULT '',
    name_ru VARCHAR(20) DEFAULT '',
    name_en VARCHAR(20) DEFAULT '',
    description_uz VARCHAR(500) DEFAULT '',
    description_ru VARCHAR(500) DEFAULT '',
    description_en VARCHAR(500) DEFAULT '',
    active BOOLEAN DEFAULT TRUE,
    order_no INTEGER DEFAULT 0,
    in_price FLOAT NOT NULL,
    out_price FLOAT NOT NULL,
    left_count INTEGER NOT NULL,
    discount_percent FLOAT DEFAULT 0,
    image VARCHAR(200)[] NOT NULL, -- array note
    created_at TIMESTAMPT DEFAULT NOW(),
    updated_at TIMESTAMPT DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE product_categories (
    id UUID PRIMARY KEY,
    product_id UUID REFERENCES product(id),
    category_id UUID REFERENCES category(id)
);

CREATE TABLE product_reviews (
    id UUID PRIMARY KEY,
    customer_id UUID NOT NULL,
    product_id UUID REFERENCES product(id),
    text VARCHAR(500) NOT NULL,
    rating FLOAT NOT NULL,
    order_id UUID NOT NULL,
    created_at TIMESTAMPT DEFAULT NOW()
);

CREATE TABLE customers (
    id UUID PRIMARY KEY,
    firstname VARCHAR(55) NOT NUll,
    lastname VARCHAR(55),
    phone VARCHAR(20) NOT NULL,
    gmail VARCHAR(30) NOT NULL,
    language VARCHAR(2) CHECK (language IN ('uz', 'ru', 'en')) NOT NULL,
    date_of_birth DATE NOT NULL,
    gender VARCHAR(6) CHECK (gender IN ('male', 'female')) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE seller (
    id UUID PRIMARY KEY,
    phone VARCHAR NOT NULL,
    gmail VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    shop_id UUID REFERENCES shop(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE system_user (
    id UUID PRIMARY KEY,
    phone VARCHAR NOT NULL,
    gmail VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    role VARCHAR(7) CHECK (role IN ('admin', 'courier')) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE branch (
    id UUID PRIMARY KEY,
    phone VARCHAR NOT NULL,
    name VARCHAR(20) DEFAULT '',
    location POLYGON NOT NULL,
    address VARCHAR NOT NULL,
    open_time TIME NOT NULL,
    close_time TIME NOT NULL,
    active BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE shop (
    id UUID PRIMARY KEY,
    slug VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    name_uz VARCHAR(20) DEFAULT '',
    name_ru VARCHAR(20) DEFAULT '',
    name_en VARCHAR(20) DEFAULT '',
    description_uz VARCHAR(500) DEFAULT '',
    description_ru VARCHAR(500) DEFAULT '',
    description_en VARCHAR(500) DEFAULT '',
    location VARCHAR NOT NULL,
    currency VARCHAR(3) CHECK (currency IN ('USD', 'EUR', 'UZS')) NOT NULL,
    payment_types VARCHAR[] NOT NULL, -- array note
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE orders (
    id UUID PRIMARY KEY,
    external_id VARCHAR NOT NULL,
    type VARCHAR(11) CHECK (type IN ('self_pickup', 'delivery')) NOT NULL,
    customer_phone VARCHAR(20) NOT NULL,
    customer_name VARCHAR(20) NOT NULL,
    customer_id UUID REFERENCES customers(id),
    payment_type VARCHAR(8) CHECK (payment_type IN ('uzum', 'cash', 'terminal')) NOT NULL,
    status VARCHAR(18) CHECK (status IN ('waiting_for_payment', 'collecting', 'delivery', 'waiting_on_branch', 'finished', 'cancelled')) NOT NULL,
    to_address VARCHAR NOT NULL,
    to_location POLYGON NOT NULL,
    discount_amount FLOAT NOT NULL,
    amount FLOAT NOT NULL,
    delivery_price FLOAT NOT NULL,
    paid BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE order_products (
    id UUID PRIMARY KEY,
    product_id UUID REFERENCES product(id),
    count INTEGER NOT NULL,
    discount_price FLOAT NOT NULL,
    price FLOAT NOT NULL,
    order_id UUID REFERENCES orders(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE order_status_notes (
    id UUID PRIMARY KEY,
    order_id UUID REFERENCES orders(id),
    status VARCHAR(18) CHECK (status IN ('waiting_for_payment', 'collecting', 'delivery', 'waiting_on_branch', 'finished', 'cancelled')) NOT NULL,
    user_id UUID NOT NULL,
    reason VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
