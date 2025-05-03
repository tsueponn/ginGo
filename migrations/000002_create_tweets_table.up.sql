CREATE TABLE IF NOT EXISTS tweets (
                        id SERIAL PRIMARY KEY,
                        content TEXT NOT NULL,
                        user_id INTEGER,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        deleted_at TIMESTAMP,

                        CONSTRAINT fk_user
                            FOREIGN KEY (user_id)
                                REFERENCES users(id)
                                ON UPDATE CASCADE
                                ON DELETE SET NULL
);