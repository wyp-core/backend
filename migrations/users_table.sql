-- Create extension for UUID support if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create enum type for gender
CREATE TYPE gender_type AS ENUM ('male', 'female', 'other');

-- Create users table with UUID primary key
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    country_code VARCHAR(5) NOT NULL,
    age INTEGER NOT NULL,
    gender gender_type NOT NULL,
    created_at TIMESTAMP 
);