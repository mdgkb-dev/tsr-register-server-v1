create table disability
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_899e53753996a8bbbf1370d2f1a"
        primary key,
    patient_id uuid not null
        constraint "FK_41eae0e7955f29a8c3bf635497c"
        references patients
        on delete cascade,
    period_id uuid not null
);

create index "IDX_41eae0e7955f29a8c3bf635497"
    on disability (patient_id);

alter table disability
    add constraint disability_period_id_fk
        foreign key (period_id) references period
            on update cascade on delete cascade;

