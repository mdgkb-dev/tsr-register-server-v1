create table anthropometry_data
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_8ef84834131dd9500d97ccf9c28"
        primary key,
    anthropometry_id uuid not null
        constraint "FK_df9d8bada0969f2ba4e1463c5db"
        references anthropometry,
    patient_id uuid not null
        constraint "FK_513cd72655542cf5ab563e9cbe9"
        references patient
        on delete cascade,
    value integer not null,
    date date not null
);

alter table anthropometry_data owner to mdgkb;

create index "IDX_df9d8bada0969f2ba4e1463c5d"
    on anthropometry_data (anthropometry_id);

create index "IDX_513cd72655542cf5ab563e9cbe"
    on anthropometry_data (patient_id);

