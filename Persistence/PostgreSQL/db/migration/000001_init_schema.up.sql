CREATE TABLE "person" (
    "id" bigserial PRIMARY KEY,
    "kind" varchar NOT NULL,
    "persons_name" varchar NOT NULL,
    "origins" varchar NOT NULL,
    "programming_languages" varchar[] NOT NULL,
    "tools" varchar[] NOT NULL,
    "linkedin" varchar NOT NULL,
    "github" varchar NOT NULL,
    "personal" varchar NOT NULL,
    "foreign_languages" varchar[] NOT NULL,
    "fav_food" varchar NOT NULL,
    "fav_drink" varchar NOT NULL,
    "fav_programming_language" varchar Not NULL,
    "thinking_about" varchar NOT NULL,
    "hobbies" varchar NOT NULL
);
