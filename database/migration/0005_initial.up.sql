CREATE TABLE IF NOT EXISTS sessions(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    expiry_time TIMESTAMP ,
    archived_at TIMESTAMP WITH TIME ZONE,
    user_id uuid REFERENCES users(id)

)