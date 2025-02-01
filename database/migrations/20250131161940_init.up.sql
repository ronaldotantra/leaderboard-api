CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "match" (
    id BIGSERIAL PRIMARY KEY,
    participant_id BIGINT[] NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "point" (
    id BIGSERIAL PRIMARY KEY,
    match_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    point INT NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "point" ADD CONSTRAINT "fk_match_id" FOREIGN KEY ("match_id") REFERENCES "match" ("id");
ALTER TABLE "point" ADD CONSTRAINT "fk_user_id" FOREIGN KEY ("user_id") REFERENCES "user" ("id");

CREATE INDEX "idx_point_date_id" ON "point" ("date");
CREATE INDEX "idx_point_user_id" ON "point" ("user_id");