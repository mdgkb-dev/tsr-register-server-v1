alter table patients
    add deleted_at timestamptz default NULL;
