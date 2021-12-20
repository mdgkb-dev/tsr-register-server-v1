alter table patients
    add region_id uuid references regions on delete cascade;
