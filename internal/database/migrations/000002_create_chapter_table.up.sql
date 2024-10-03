CREATE TABLE IF NOT EXISTS chapters (
    chapter_id uuid DEFAULT gen_random_uuid(),
    novel_id uuid NOT NULL,
    title VARCHAR(255) NOT NULL,
    content VARCHAR(255) NOT NULL,

    PRIMARY KEY(novel_id),
    CONSTRAINT fk_novel
        FOREIGN KEY(novel_id)
            REFERENCES novels(novel_id)
) 