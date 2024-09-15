-- +goose Up
-- +goose StatementBegin

/* deprecated */
CREATE TABLE permits (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  permitID TEXT ,
  companyName TEXT ,
  reference TEXT,
  dateReceived DATETIME,
  dateDue DATETIME,
  permitStatus TEXT,
  designer TEXT
);
/* deprecated */

CREATE TABLE permitid (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  permitID TEXT
);

CREATE TABLE permit_company (
  permit TEXT,
  companyName TEXT,
  FOREIGN KEY (permit) REFERENCES permitid(id) NOT NULL
);

CREATE TABLE permit_designer (
  permit TEXT,
  designer TEXT,
  dateStarted DATETIME,
  dateCompleted DATETIME, 
  FOREIGN KEY (permit) REFERENCES permitid(id) NOT NULL
);

CREATE TABLE permit_date_received (
  permit TEXT,
  dateReceived DATETIME,
  FOREIGN KEY (permit) REFERENCES permitid(id) NOT NULL
);

CREATE TABLE permit_date_due (
  permit TEXT,
  dateDue DATETIME
  FOREIGN KEY (permit) REFERENCES permitid(id) NOT NULL
);

CREATE TABLE permit_date_submited (
  permit TEXT,
  dateSubmited DATETIME
  FOREIGN KEY (permit) REFERENCES permitid(id) NOT NULL
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE permits;
-- +goose StatementEnd