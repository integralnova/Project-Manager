-- +goose Up
-- +goose StatementBegin

CREATE TABLE permitid (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  permitID TEXT
);

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

CREATE TABLE permit_company (
  permit INTEGER,
  companyName TEXT
);

CREATE TABLE permit_designer (
  permit INTEGER,
  designer TEXT,
  dateStarted DATETIME,
  dateCompleted DATETIME
);

CREATE TABLE permit_date_received (
  permit INTEGER,
  dateReceived DATETIME
);

CREATE TABLE permit_date_due (
  permit INTEGER,
  dateDue DATETIME
);

CREATE TABLE permit_date_submited (
  permit INTEGER,
  dateSubmited DATETIME
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE permits;
DROP TABLE permitid;
DROP TABLE permit_company;
DROP TABLE permit_designer;
DROP TABLE permit_date_received;
DROP TABLE permit_date_due;
DROP TABLE permit_date_submited;
-- +goose StatementEnd
