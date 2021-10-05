alter table patient
    add deleted_at timestamptz default NULL;
