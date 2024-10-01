CREATE TABLE IF NOT EXISTS novels (
 id uuid DEFAULT gen_random_uuid(),
 author_id uuid NOT NULL, 
 title VARCHAR(255) NOT NULL,
 synopsis VARCHAR(255) NOT NULL,
 rating DOUBLE PRECISION DEFAULT 0,
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO novels (author_id, title, synopsis, rating)
VALUES 
  ('d9b08e9d-bd2e-4c32-9e7a-7c8b51b6e015', 'The Lost Chronicles', 'An epic journey through forgotten realms.', 4.5),
  ('f7a1d23e-8af2-432e-9bdf-2e3bda5d2c1f', 'Whispers in the Dark', 'A gripping tale of mystery and suspense.', 4.8),
  ('2e5d7c2e-b2b4-4f16-918a-121d91aeaad1', 'Beneath the Stars', 'A love story set in the vastness of space.', 4.2),
  ('5c8e5f4b-6bf5-4899-9084-9d92bc7a0f7f', 'Shadows of the Forgotten', 'A thrilling adventure filled with dark secrets.', 4.6);
