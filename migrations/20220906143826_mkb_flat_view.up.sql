create view mkb_flat_view(id, full_name, level, class_id) as
SELECT mc.id,
       concat(mc.range_start, '-', mc.range_end, ' ', mc.number, ' ', mc.name) AS full_name,
       0                                                                       AS level,
        mc.id as class_id
FROM mkb_class mc
UNION ALL
SELECT mg.id,
       concat(mg.range_start, '-', mg.range_end, ' ', mg.name) AS full_name,
       1                                                       AS level,
mg.mkb_class_id as class_id
FROM mkb_groups mg
UNION ALL
SELECT msg.id,
       concat(msg.range_start, '-', msg.range_end, ' ', msg.name) AS full_name,
       2                                                          AS level,
g.mkb_class_id as class_id
FROM mkb_sub_group msg
join mkb_groups g on msg.mkb_group_id = g.id
UNION ALL
SELECT mssg.id,
       concat(mssg.range_start, '-', mssg.range_end, ' ', mssg.name) AS full_name,
       3                                                             AS level,
mkb_groups.mkb_class_id as class_id
FROM mkb_sub_sub_group mssg
join mkb_sub_group  on mkb_sub_group.id = mssg.mkb_sub_group_id
join mkb_groups  on  mkb_sub_group.mkb_group_id = mkb_groups.id
UNION ALL
SELECT md.id,
       concat(md.code, ' ', md.name) AS full_name,
       4                             AS level,
md.mkb_class_id as class_id
FROM mkb_diagnosis md
UNION ALL
SELECT msd.id,
       concat(m.code, '.', msd.sub_code, ' ', msd.name) AS full_name,
       5                                                AS level,
       m.mkb_class_id as class_id
FROM mkb_sub_diagnosis msd
         JOIN mkb_diagnosis m ON m.id = msd.mkb_diagnosis_id
UNION ALL
SELECT mcd.id,
       concat(mcd.name) AS full_name,
       6                AS level,
       md.mkb_class_id as class_id
FROM mkb_concrete_diagnosis mcd
join mkb_sub_diagnosis msd on mcd.mkb_sub_diagnosis_id = msd.id
join mkb_diagnosis md on msd.mkb_diagnosis_id = md.id;


