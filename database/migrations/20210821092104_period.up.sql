CREATE TABLE period
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    date_start date not null ,
    date_end date  not null
);
