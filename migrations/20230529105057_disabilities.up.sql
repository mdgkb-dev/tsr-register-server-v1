alter table disabilities add column date_start date;
alter table disabilities add column date_end date;

update disabilities
set date_start = (select p.date_start from period p where p.id = disabilities.period_id)
where period_id is not null;

update disabilities
set date_end = (select p.date_end from period p where p.id = disabilities.period_id)
where period_id is not null;

alter table disabilities drop column period_id;

alter table edvs add column date_start date;
alter table edvs add column date_end date;
alter table edvs drop column period_id;