create view mkb_groups_view as
SELECT
    m.*,
    concat(m.range_start, '-', m.range_end, ' ', m.name) as full_name
FROM mkb_groups m;

create view mkb_diagnosis_view as
select m.*, concat(m.code, ' ', m.name) as full_name
from mkb_diagnosis m;


create view mkb_sub_diagnosis_view as
select msd.*, concat(md.code, '.', msd.sub_code, ' ', msd.name) as full_name
from mkb_sub_diagnosis msd
join mkb_diagnosis md on msd.mkb_diagnosis_id = md.id;