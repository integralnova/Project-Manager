-- +goose Up
-- +goose StatementBegin


CREATE TABLE permits (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  permitID TEXT,
  companyName TEXT,
  reference TEXT,
  dateReceived DATETIME,
  dateDue DATETIME,
  permitStatus TEXT,
  designer TEXT
);




-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE permits;
-- +goose StatementEnd
