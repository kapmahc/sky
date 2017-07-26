-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE survey_forms (
  id         BIGSERIAL PRIMARY KEY,
  title      VARCHAR(255)                NOT NULL,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  deadline   DATE NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_survey_forms_type
  ON survey_forms (type);

CREATE TABLE survey_fields (
  id         BIGSERIAL PRIMARY KEY,
  label      VARCHAR(255)                NOT NULL,
  name       VARCHAR(255)                NOT NULL,
  value      VARCHAR(255)                NOT NULL,
  body       TEXT                        NOT NULL,
  type       VARCHAR(16)                 NOT NULL DEFAULT 'text',
  required   BOOLEAN                     NOT NULL DEFAULT TRUE,
  form_id    BIGINT                      REFERENCES survey_forms,
  sort_order INT                         NOT NULL DEFAULT 0,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_survey_fields_name_form_id
  ON survey_fields (name, form_id);

CREATE TABLE survey_records (
  id         BIGSERIAL PRIMARY KEY,
  username   VARCHAR(255)                NOT NULL,
  email      VARCHAR(255)                NOT NULL,
  phone      VARCHAR(255)                NOT NULL,
  value      TEXT                        NOT NULL,
  form_id    BIGINT                      REFERENCES survey_forms,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_survey_fields_email_form
  ON survey_records (email, form_id);
CREATE UNIQUE INDEX idx_survey_fields_phone_phone
  ON survey_records (phone, form_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE survey_records;
DROP TABLE survey_fields;
DROP TABLE survey_forms;
