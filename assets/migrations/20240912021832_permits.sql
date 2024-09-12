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

INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240228', 'kgcdt', 'CL', '6/23/2010', '11/9/2005', 'QC', 'Shelby Rijkeseis');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240629', 'uyxnl', 'BH', '2/4/2008', '5/27/2007', 'Pole Foreman', 'Wendall Tidcomb');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240703', 'omhvq', 'BH', '12/5/2001', '7/27/2006', 'WO', 'Tami Batrip');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240114', 'qohyv', 'CL', '1/20/2007', '6/15/2016', 'Design', 'Ingeborg Leake');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240203', 'qylbl', 'BH', '12/21/2012', '7/19/2009', 'QC', 'Dulce MacDunleavy');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240765', 'fvaxj', 'ATT', '6/24/2014', '2/4/2020', 'Design', 'Thorstein Wasiel');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240370', 'qdbxb', 'ATT', '8/1/2013', '10/26/2003', 'Pole Foreman', 'Ethelin MacFaul');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240498', 'tuyjn', 'CL', '9/20/2022', '10/31/2002', 'Design', 'Helyn Lackemann');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240680', 'pelrc', 'BH', '5/20/2017', '2/16/2022', '', 'Mendie Dekeyser');
INSERT INTO permits (permitID, reference, companyName, dateReceived, dateDue, permitStatus, designer) VALUES ('240896', 'vglcn', 'BH', '9/22/2015', '7/21/2013', 'WO', 'Owen Lody');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE permits;
-- +goose StatementEnd