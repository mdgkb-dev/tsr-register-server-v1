create table commissions_statuses
(
    id   uuid default uuid_generate_v4() not null primary key,
    name varchar,
    color varchar
);

alter table commissions add column commission_status_id uuid REFERENCES commissions_statuses(id) ON UPDATE CASCADE ON DELETE CASCADE;
alter table commissions add column dzm_answer_file_id uuid REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE;
alter table commissions add column dzm_answer_comment varchar;

insert into commissions_statuses (id, name, color)
values  ('9863d5b6-166f-459c-aadc-bb594d10993b', 'Создана', '#536AC2'),
        ('ef85a78d-fe41-461d-b2d4-7063389fd055', 'Проведена', '#BFB467'),
        ('9f1f4e77-bf4e-47fc-a515-7092285fa69e', 'На рассмотрении', '#92D2D0'),
        ('f8c8dd10-93ff-4ebb-af05-1b408e8b0319', 'Отказ', '#FF0000');


