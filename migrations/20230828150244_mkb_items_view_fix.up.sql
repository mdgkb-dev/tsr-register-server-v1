drop view mkb_items_view;

create view mkb_items_view(id, full_name) as
SELECT mi.id,
       CASE
           WHEN mi.code::text = '-'::text THEN concat(mi.range_start, '-', mi.range_end, ' ', mi.name)
           WHEN mi.range_start::text = '-'::text AND mi.range_end::text = '-'::text THEN concat(mi.code, ' ', mi.name)
           ELSE NULL::text
           END AS full_name
FROM mkb_items mi;
