CREATE TYPE mode_type AS ENUM ('remote', 'onsite', 'hybrid');

CREATE TABLE jobs (
    job_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_by UUID NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    lat DOUBLE PRECISION NOT NULL,
    lon DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    price DOUBLE PRECISION NOT NULL,
    category TEXT,
    mode mode_type NOT NULL,
    views INTEGER DEFAULT 0,
    duration TEXT,
    geo_location GEOMETRY ,
);