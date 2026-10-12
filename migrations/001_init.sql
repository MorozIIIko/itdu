-- +goose Up
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    telegram_id BIGINT NOT NULL UNIQUE,
    username text ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE TABLE quizzes (
id BIGSERIAL PRIMARY KEY,
name text NOT NULL
);
CREATE TABLE questions (
id BIGSERIAL PRIMARY KEY,
text text NOT NULL,
quiz_id BIGINT NOT NULL REFERENCES quizzes(id)
);
CREATE TABLE answer_options (
id BIGSERIAL PRIMARY KEY, 
text text NOT NULL,
question_id BIGINT NOT NULL REFERENCES questions(id),
is_correct BOOLEAN NOT NULL DEFAULT FALSE
);
CREATE TABLE quiz_attempts (
id BIGSERIAL PRIMARY KEY, 
user_id BIGINT NOT NULL REFERENCES users(id),
quiz_id BIGINT NOT NULL REFERENCES quizzes(id),
started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
finished_at TIMESTAMPTZ
);
CREATE TABLE user_answers (
id BIGSERIAL PRIMARY KEY,
attempt_id BIGINT NOT NULL REFERENCES quiz_attempts(id),
question_id BIGINT NOT NULL REFERENCES questions(id),
selected_option_id BIGINT NOT NULL REFERENCES answer_options(id),
answered_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
UNIQUE(attempt_id, question_id)
);
-- +goose Down
DROP TABLE user_answers;
DROP TABLE quiz_attempts;
DROP TABLE answer_options;
DROP TABLE questions;
DROP TABLE quizzes;
DROP TABLE users;