-- Initial migration 
-- Database: PostgreSQL

BEGIN;

-- Create ENUM types
CREATE TYPE user_role AS ENUM ('customer', 'admin');
CREATE TYPE address_type AS ENUM ('shipping', 'billing');
CREATE TYPE order_status AS ENUM ('pending', 'paid', 'shipped', 'delivered', 'cancelled');
CREATE TYPE payment_method AS ENUM ('telebirr', 'cbe_banking');
CREATE TYPE auth_provider AS ENUM ('local', 'google');

-- 1. Users Table
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

-- 2. Addresses Table
CREATE TABLE addresses (
    address_id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    street VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100),
    postal_code VARCHAR(20) NOT NULL,
    country VARCHAR(100) NOT NULL,
    address_type address_type NOT NULL
);

-- 3. Categories Table
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(100) NOT NULL,
    parent_category_id INTEGER REFERENCES categories(category_id) ON DELETE SET NULL
);

-- 4. Brands Table
CREATE TABLE brands (
    brand_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);

-- 5. Products Table
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL REFERENCES categories(category_id) ON DELETE RESTRICT,
    brand_id INTEGER NOT NULL REFERENCES brands(brand_id) ON DELETE RESTRICT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock_quantity INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT positive_price CHECK (price >= 0),
    CONSTRAINT positive_stock CHECK (stock_quantity >= 0)
);

-- 6. Product_Specifications Table (Composite PK)
CREATE TABLE product_specifications (
    product_id INTEGER NOT NULL,
    spec_name VARCHAR(100) NOT NULL,
    spec_value VARCHAR(255) NOT NULL,
    PRIMARY KEY (product_id, spec_name),
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- 7. Compatibility_Rules Table
CREATE TABLE compatibility_rules (
    product_id INTEGER NOT NULL REFERENCES products(product_id) ON DELETE CASCADE,
    spec_id INTEGER NOT NULL REFERENCES product_specifications(product_id, spec_name) ON DELETE CASCADE,
    spec_value VARCHAR(255) NOT NULL,
    PRIMARY KEY (product_id, spec_id, spec_value)
);

-- 8. Custom_Builds Table
CREATE TABLE custom_builds (
    build_id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE SET NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    total_price DECIMAL(10,2) NOT NULL DEFAULT 0,
    CONSTRAINT positive_total_price CHECK (total_price >= 0)
);

-- 9. Build_Items Table (Composite PK)
CREATE TABLE build_items (
    build_id UUID NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (build_id, product_id),
    FOREIGN KEY (build_id) REFERENCES custom_builds(build_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE RESTRICT,
    CONSTRAINT positive_quantity CHECK (quantity > 0)
);

-- 10. Orders Table (Merged with Payments)
CREATE TABLE orders (
    order_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    address_id INTEGER NOT NULL REFERENCES addresses(address_id) ON DELETE RESTRICT,
    order_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status order_status NOT NULL DEFAULT 'pending',
    total_amount DECIMAL(10,2) NOT NULL,
    payment_method payment_method NOT NULL,
    payment_date TIMESTAMP,
    CONSTRAINT positive_total_amount CHECK (total_amount >= 0)
);

-- 11. Order_Items Table (Composite PK)
CREATE TABLE order_items (
    order_id UUID NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders(order_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE RESTRICT,
    CONSTRAINT positive_quantity CHECK (quantity > 0),
    CONSTRAINT positive_unit_price CHECK (unit_price >= 0)
);

-- 12. Reviews Table
CREATE TABLE reviews (
    review_id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    product_id INTEGER NOT NULL REFERENCES products(product_id) ON DELETE CASCADE,
    rating INTEGER NOT NULL CHECK (rating BETWEEN 1 AND 5),
    comment TEXT,
    review_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 13. Cart Table (Composite PK)
CREATE TABLE cart (
    user_id UUID NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    PRIMARY KEY (user_id, product_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE,
    CONSTRAINT positive_quantity CHECK (quantity > 0)
);

-- 14. Suppliers Table
CREATE TABLE suppliers (
    supplier_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    contact_email VARCHAR(255),
    phone VARCHAR(20)
);

-- 15. Inventory Table
CREATE TABLE inventory (
    inventory_id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(product_id) ON DELETE CASCADE,
    supplier_id INTEGER NOT NULL REFERENCES suppliers(supplier_id) ON DELETE RESTRICT,
    quantity INTEGER NOT NULL,
    last_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT positive_quantity CHECK (quantity >= 0)
);

-- 16. Discounts Table
CREATE TABLE discounts (
    discount_id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES products(product_id) ON DELETE SET NULL,
    category_id INTEGER REFERENCES categories(category_id) ON DELETE SET NULL,
    discount_percentage DECIMAL(5,2) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    CONSTRAINT valid_percentage CHECK (discount_percentage BETWEEN 0 AND 100),
    CONSTRAINT valid_dates CHECK (start_date <= end_date)
);

-- Indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_provider_id ON users(provider, provider_id);
CREATE INDEX idx_addresses_user_id ON addresses(user_id);
CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_brand_id ON products(category_id, brand_id);
CREATE INDEX idx_product_specifications_product_id ON product_specifications(product_id);
CREATE INDEX idx_compatibility_rules_product_ids ON compatibility_rules(product_id_1, product_id_2);
CREATE INDEX idx_custom_builds_user_id ON custom_builds(user_id);
CREATE INDEX idx_build_items_build_id ON build_items(build_id);
CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orders_payment_date ON orders(payment_date);
CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_reviews_product_id ON reviews(product_id);
CREATE INDEX idx_cart_user_id ON cart(user_id);
CREATE INDEX idx_inventory_product_id ON inventory(product_id);
CREATE INDEX idx_discounts_product_id ON discounts(product_id);

-- Trigger Function: Update stock_quantity in products based on inventory
CREATE OR REPLACE FUNCTION update_product_stock()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE products
    SET stock_quantity = (
        SELECT COALESCE(SUM(quantity), 0)
        FROM inventory
        WHERE product_id = NEW.product_id
    )
    WHERE product_id = NEW.product_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER inventory_update_trigger
AFTER INSERT OR UPDATE OR DELETE ON inventory
FOR EACH ROW
EXECUTE FUNCTION update_product_stock();

-- Trigger Function: Update total_price in custom_builds
CREATE OR REPLACE FUNCTION update_build_total_price()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE custom_builds
    SET total_price = (
        SELECT COALESCE(SUM(p.price * bi.quantity), 0)
        FROM build_items bi
        JOIN products p ON bi.product_id = p.product_id
        WHERE bi.build_id = NEW.build_id
    )
    WHERE build_id = NEW.build_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER build_items_update_trigger
AFTER INSERT OR UPDATE OR DELETE ON build_items
FOR EACH ROW
EXECUTE FUNCTION update_build_total_price();

-- Stored Procedure: Place Order
CREATE OR REPLACE PROCEDURE place_order(
    p_user_id UUID,
    p_address_id INTEGER,
    p_payment_method payment_method,
    p_items JSONB
)
LANGUAGE plpgsql AS $$
DECLARE
    v_order_id UUID;
    v_item JSONB;
BEGIN
    -- Insert order
    INSERT INTO orders (user_id, address_id, payment_method, total_amount)
    VALUES (
        p_user_id,
        p_address_id,
        p_payment_method,
        (SELECT SUM((item->>'quantity')::INTEGER * p.price)
         FROM JSONB_ARRAY_ELEMENTS(p_items) item
         JOIN products p ON (item->>'product_id')::INTEGER = p.product_id)
    )
    RETURNING order_id INTO v_order_id;

    -- Insert order items
    FOR v_item IN SELECT * FROM JSONB_ARRAY_ELEMENTS(p_items)
    LOOP
        INSERT INTO order_items (order_id, product_id, quantity, unit_price)
        SELECT
            v_order_id,
            (v_item->>'product_id')::INTEGER,
            (v_item->>'quantity')::INTEGER,
            p.price
        FROM products p
        WHERE p.product_id = (v_item->>'product_id')::INTEGER;

        -- Update inventory
        UPDATE inventory
        SET quantity = quantity - (v_item->>'quantity')::INTEGER
        WHERE product_id = (v_item->>'product_id')::INTEGER
        AND quantity >= (v_item->>'quantity')::INTEGER;
    END LOOP;

    COMMIT;
EXCEPTION
    WHEN OTHERS THEN
        ROLLBACK;
        RAISE EXCEPTION 'Order placement failed: %', SQLERRM;
END;
$$;

-- Stored Procedure: Validate Build Compatibility
CREATE OR REPLACE FUNCTION validate_build(p_build_id UUID)
RETURNS TABLE (is_compatible BOOLEAN, message TEXT) AS $$
DECLARE
    v_item RECORD;
    v_rule RECORD;
BEGIN
    is_compatible := TRUE;
    message := '';

    FOR v_item IN
        SELECT bi.product_id
        FROM build_items bi
        WHERE bi.build_id = p_build_id
    LOOP
        FOR v_rule IN
            SELECT cr.spec_id, ps.spec_name, cr.spec_value
            FROM compatibility_rules cr
            JOIN product_specifications ps ON cr.product_id = ps.product_id AND cr.spec_id = ps.spec_name
            WHERE cr.product_id = v_item.product_id
            AND cr.spec_value != 'ANY' -- Skip universal compatibility rules 
        LOOP
            IF NOT EXISTS (
                SELECT 1
                FROM build_items bi2
                JOIN product_specifications ps2 ON bi2.product_id = ps2.product_id
                WHERE bi2.build_id = p_build_id
                AND ps2.spec_name = v_rule.spec_name
                AND ps2.spec_value = v_rule.spec_value
            ) THEN
                is_compatible := FALSE;
                message := message || format('Incompatible: Product %s requires %s=%s; ', 
                                            v_item.product_id, v_rule.spec_name, v_rule.spec_value);
            END IF;
        END LOOP;
    END LOOP;

    RETURN NEXT;
END;
$$ LANGUAGE plpgsql;

-- Stored Function: Get Compatible Products
CREATE OR REPLACE FUNCTION get_compatible_products(p_build_id UUID, p_category_name TEXT)
RETURNS TABLE (product_id INTEGER, product_name TEXT) AS $$
BEGIN
    RETURN QUERY
    SELECT p.product_id, p.name AS product_name
    FROM products p
    JOIN categories c ON p.category_id = c.category_id
    WHERE c.name = p_category_name
    AND NOT EXISTS (
        -- Check for any incompatible rules
        SELECT 1
        FROM compatibility_rules cr
        JOIN product_specifications ps ON cr.product_id = ps.product_id AND cr.spec_id = ps.spec_name
        WHERE cr.product_id = p.product_id
        AND cr.spec_value != 'ANY' -- Skip universal compatibility rules (e.g., RAM)
        AND NOT EXISTS (
            -- Verify selected components meet the rule
            SELECT 1
            FROM build_items bi
            JOIN product_specifications ps2 ON bi.product_id = ps2.product_id
            WHERE bi.build_id = p_build_id
            AND ps2.spec_name = ps.spec_name
            AND ps2.spec_value = cr.spec_value
        )
    );
END;
$$ LANGUAGE plpgsql;

-- Views
CREATE VIEW top_rated_products AS
SELECT p.product_id, p.name, AVG(r.rating) AS avg_rating
FROM products p
JOIN reviews r ON p.product_id = r.product_id
GROUP BY p.product_id, p.name
HAVING COUNT(r.review_id) > 0;

CREATE VIEW user_cart_summary AS
SELECT c.user_id, p.product_id, p.name, c.quantity, p.price, (c.quantity * p.price) AS total
FROM cart c
JOIN products p ON c.product_id = p.product_id;

COMMIT;
