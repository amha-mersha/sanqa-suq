BEGIN; 
-- ENUM 
CREATE TYPE user_role AS ENUM('customer', 'seller');
CREATE TYPE address_type AS ENUM('billing', 'shipping');
CREATE TYPE order_status AS ENUM('shipping', 'pending', 'paid', 'delivered', 'cancelled');
CREATE TYPE payment_method AS ENUM('cbe', 'telebirr');
CREATE TYPE address_type AS ENUM('local', 'google');

-- User
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone VARCHAR(20),
    role user_role NOT NULL DEFAULT 'customer',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    provider auth_provider NOT NULL DEFAULT 'local',
    provider_id VARCHAR(255),
    CONSTRAINT unique_provider_id CHECK (provider = 'local' OR provider_id IS NOT NULL)
);

-- Address 
CREATE TABLE addresses (
    address_id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    street VARCHAR(255),
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100),
    postal_code VARCHAR(20) NOT NULL,
    country VARCHAR(100) NOT NULL,
    address_type address_type NOT NULL
);

-- categories
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    parent_category_id INTEGER REFERENCES categories(category_id) ON DELETE SET NULL
);
