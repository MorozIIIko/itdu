-- +goose Up

INSERT INTO quizzes (name)
VALUES
    ('Go'),
    ('SQL');

-- +goose Down

DELETE FROM quizzes
WHERE name IN ('Go', 'SQL');