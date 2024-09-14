-- +goose Up
-- +goose StatementBegin
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

CREATE TABLE permitid (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  permitID TEXT
);

CREATE TABLE permit_company (
  permit TEXT,
  companyName TEXT
);

CREATE TABLE permit_designer (
  permit TEXT,
  designer TEXT,
  dateStarted DATETIME,
  dateCompleted DATETIME //refactor need goose up
);

CREATE TABLE permit_date_received (
  permit TEXT,
  dateReceived DATETIME
);

CREATE TABLE permit_date_due (
  permit TEXT,
  dateDue DATETIME
);

CREATE TABLE permit_date_submited (
  permit TEXT,
  dateSubmited DATETIME
);

INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240228', 'kgcdt', 'CL', '2010-06-23 00:00:00', '2005-11-09 00:00:00', 'QC', 'Shelby Rijkeseis');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240629', 'uyxnl', 'BH', '2008-02-04 00:00:00', '2007-05-27 00:00:00', 'Pole Foreman', 'Wendall Tidcomb');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240703', 'omhvq', 'BH', '2001-12-05 00:00:00', '2006-07-27 00:00:00', 'WO', 'Tami Batrip');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240114', 'qohyv', 'CL', '2007-01-20 00:00:00', '2016-06-15 00:00:00', 'Design', 'Ingeborg Leake');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240815', 'abxyt', 'CL', '2012-03-14 00:00:00', '2013-08-22 00:00:00', 'Approved', 'Jordan Smith');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240916', 'cdzuv', 'BH', '2015-07-19 00:00:00', '2015-12-30 00:00:00', 'Pending', 'Alex Johnson');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('241017', 'efwxy', 'CL', '2018-11-11 00:00:00', '2019-04-05 00:00:00', 'Rejected', 'Taylor Brown');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('241118', 'ghyzb', 'BH', '2020-05-23 00:00:00', '2020-10-15 00:00:00', 'In Review', 'Morgan Lee');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('241219', 'ijklm', 'CL', '2021-09-30 00:00:00', '2022-02-20 00:00:00', 'Completed', 'Casey Davis');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('241320', 'mnopq', 'BH', '2023-01-10 00:00:00', '2023-06-25 00:00:00', 'Under Review', 'Jamie Clark');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('241421', 'rstuv', 'CL', '2019-04-18 00:00:00', '2019-09-12 00:00:00', 'Approved', 'Riley Morgan');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('241522', 'wxyza', 'BH', '2017-08-05 00:00:00', '2018-01-20 00:00:00', 'Pending', 'Jordan Taylor');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('241623', 'bcdef', 'CL', '2014-12-15 00:00:00', '2015-05-10 00:00:00', 'Rejected', 'Casey Lee');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('241724', 'ghijk', 'BH', '2011-06-22 00:00:00', '2011-11-30 00:00:00', 'Completed', 'Morgan Davis');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE permits;
-- +goose StatementEnd