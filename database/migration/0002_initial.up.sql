CREATE TABLE IF NOT EXISTS tasks(
    id  uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL ,
    created_at TIMESTAMP WITH TIME ZONE,
    is_completed BOOL DEFAULT FALSE,
    archived_at TIMESTAMP WITH TIME ZONE,
    user_id UUID REFERENCES users(id)
);