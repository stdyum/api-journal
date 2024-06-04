-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS journal.lessons_info
(
    id             uuid,
    study_place_id uuid,
    teacher_id     uuid,
    lesson_id      uuid,
    title          varchar,
    description    varchar,
    homework       varchar,
    type           varchar,
    created_at     timestamp,
    updated_at     timestamp,

    PRIMARY KEY ((study_place_id), lesson_id, id),
);

-- +goose StatementEnd
