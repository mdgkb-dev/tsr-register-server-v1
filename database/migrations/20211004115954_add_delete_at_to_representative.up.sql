alter table representative
    add deleted_at timestamptz default NULL;
