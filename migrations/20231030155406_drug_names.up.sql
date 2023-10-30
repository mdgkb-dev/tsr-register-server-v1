create or replace view drugs_view as
SELECT
    d.*,
    CONCAT_WS(' '::TEXT, d.name, d.name_inn) AS search_name
FROM drugs d;
         

create or replace view drugs_view as
SELECT
    d.*,
    CONCAT_WS(' '::TEXT, d.name, d.name_inn) AS search_name
FROM drugs d;

update search_groups
set 
search_column = 'search_name',
search_group_table = 'drugs_view',
label_column = 'search_name'
WHERE key = 'drug';
