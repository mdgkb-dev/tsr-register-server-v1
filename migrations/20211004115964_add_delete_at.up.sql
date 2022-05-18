alter table human
    add deleted_at timestamptz default NULL;

alter table human_histories
    add deleted_at timestamptz default NULL;

alter table representative_to_patient
    add deleted_at timestamptz default NULL;

alter table register_property_set_to_patient
    add deleted_at timestamptz default NULL;

alter table register_to_patient
    add deleted_at timestamptz default NULL;

alter table register_property_to_patient
    add deleted_at timestamptz default NULL;

alter table patient_drug_regimens
    add deleted_at timestamptz default NULL;

alter table patient_diagnosis
    add deleted_at timestamptz default NULL;

alter table disability
    add deleted_at timestamptz default NULL;

alter table height_weight
    add deleted_at timestamptz default NULL;

alter table contact
    add deleted_at timestamptz default NULL;

alter table file_infos
    add deleted_at timestamptz default NULL;

alter table document
    add deleted_at timestamptz default NULL;

alter table insurance_company_to_human
    add deleted_at timestamptz default NULL;
