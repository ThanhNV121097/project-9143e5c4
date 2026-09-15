CREATE TABLE IF NOT EXISTS greetings (
    id smallint PRIMARY KEY CHECK (id = 1),
    text text NOT NULL CHECK (btrim(text) <> ''),
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO greetings (id, text)
VALUES (1, 'Hello, World!')
ON CONFLICT (id) DO NOTHING;
