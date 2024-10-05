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

INSERT INTO permits (permitID, companyName, reference, dateReceived, dateDue, permitStatus, designer)
VALUES
    ('P001', 'Acme Rocketry Co.', 'Mission to Mars', '2024-10-01 09:00:00', '2024-10-15 17:00:00', 'Pending', 'Wile E. Coyote'),
    ('P002', 'Wonderland Tea Party Services', 'Mad Hatter''s Extravaganza', '2024-09-20 14:30:00', '2024-10-10 12:00:00', 'Approved', 'Alice Liddell'),
    ('P003', 'Pirate Cove Adventures', 'X Marks the Spot', '2024-09-25 11:15:00', '2024-10-05 16:30:00', 'Completed', 'Captain Hook'),
    ('P004', 'Chocolate Factory Inc.', 'Golden Ticket Production', '2024-09-15 08:00:00', '2024-09-30 18:00:00', 'Approved', 'Oompa-Loompa Engineering'),
    ('P005', 'Sherwood Forest Construction', 'Merry Men Hideout', '2024-09-18 10:30:00', '2024-10-03 14:00:00', 'Approved', 'Friar Tuck');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE permits;
-- +goose StatementEnd
