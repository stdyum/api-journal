-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS journal.absences
(
    id             uuid,
    study_place_id uuid,
    student_id     uuid,
    teacher_id     uuid,
    lesson_id      uuid,
    absence        int,
    created_at     timestamp,
    updated_at     timestamp,

    PRIMARY KEY ((study_place_id), lesson_id, id, teacher_id),
);

-- +goose StatementEnd
