ALTER TABLE search_groups add column join_table varchar;
ALTER TABLE search_groups add column join_column varchar;


update search_groups set join_table = 'patients_domains', join_column = 'patient_id' WHERE key = 'patient';