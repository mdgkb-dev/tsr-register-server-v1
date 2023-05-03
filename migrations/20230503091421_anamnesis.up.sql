alter table patient_diagnosis_anamnesis rename to anamneses;
alter table anamneses add column doctor_name varchar;
alter table anamneses rename column date to item_date;