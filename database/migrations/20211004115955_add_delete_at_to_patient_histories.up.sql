alter table patient_histories
    add deleted_at timestamptz default NULL;
