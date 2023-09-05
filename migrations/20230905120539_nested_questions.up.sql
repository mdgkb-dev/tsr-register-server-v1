alter table questions drop with_dates, drop tag, drop  col_width, drop is_files_storage, drop domain_id;

insert into questions (id, name, item_order, value_type_id, with_other,  research_id,  short_name,  age_compare,  code, calculate_scores, parent_id, comment)
select '4fe6ff2d-7528-4809-926e-32b3ebd2be71', name, item_order, value_type_id, with_other, research_id, short_name, age_compare, code, calculate_scores, parent_id, comment
from questions where id = 'd8c2be17-bace-45c1-8eda-d76f16b27a59';


UPDATE questions SET name = 'Была ли вызвана СМП?' WHERE id = 'd8c2be17-bace-45c1-8eda-d76f16b27a59';
UPDATE questions SET name = 'Какое лечение проводилось?' WHERE id = '4fe6ff2d-7528-4809-926e-32b3ebd2be71';


UPDATE public.answer_variants SET item_order = 0 WHERE id = '22f8abe1-1548-42c1-8332-a1454feea1b5';
UPDATE public.answer_variants SET item_order = 1 WHERE id = '4f3d9ad9-04fc-4dfd-9b9c-b039c8907e18';

update questions set name = 'Требовалось ли повторное введение адреналина?' where id ='01167b82-a927-4aa9-80f3-4ac3c5558aa0';


update answer_variants set show_more_questions = false
where id != '84c7ee42-3db3-4dce-9a5f-b74b16e1cf42' and question_id = '4fe6ff2d-7528-4809-926e-32b3ebd2be71';

update questions set parent_id = '4fe6ff2d-7528-4809-926e-32b3ebd2be71'
where id = '442dca02-5491-4e45-839f-c39c81ab6687';

update questions set value_type_id = 'fc00cc5a-f7a5-4974-ad57-9432656d5e0e' where  id = 'd8c2be17-bace-45c1-8eda-d76f16b27a59';

select *
from answer_variants where question_id = 'd8c2be17-bace-45c1-8eda-d76f16b27a59';

update answer_variants set question_id = '4fe6ff2d-7528-4809-926e-32b3ebd2be71'
where id in
      ('6ce4ba5f-fe5d-4418-bdda-90ec05f72e05',
       '42840878-bbf7-4e19-9281-94b58b8ca61c',
       '32f887a8-ea2e-4d56-8cad-fde09531cd76',
       '74d2f600-60a1-4dbf-b958-1d39fc1ae0f2',
       '84c7ee42-3db3-4dce-9a5f-b74b16e1cf42');


UPDATE answer_variants SET item_order = 0 WHERE id = '6ce4ba5f-fe5d-4418-bdda-90ec05f72e05';
UPDATE answer_variants SET item_order = 1 WHERE id = '42840878-bbf7-4e19-9281-94b58b8ca61c';


UPDATE questions SET item_order = 0, parent_id = 'd8c2be17-bace-45c1-8eda-d76f16b27a59', research_id = null WHERE id = '4fe6ff2d-7528-4809-926e-32b3ebd2be71';
UPDATE questions SET item_order = 1, parent_id = 'd8c2be17-bace-45c1-8eda-d76f16b27a59', research_id = null WHERE id = 'd8c2be17-bace-45c1-8eda-d76f16b27a58';
UPDATE questions SET item_order = 2, parent_id = 'd8c2be17-bace-45c1-8eda-d76f16b27a59', research_id = null  WHERE id = 'd3ff9beb-c9b7-480f-9873-7bfb71a35863';
UPDATE questions SET item_order = 3, parent_id = 'd8c2be17-bace-45c1-8eda-d76f16b27a59', research_id = null  WHERE id = 'd8c2be17-bace-45c1-8eda-d76f16b27a61';
UPDATE questions SET item_order = 4, parent_id = 'd8c2be17-bace-45c1-8eda-d76f16b27a59', research_id = null  WHERE id = 'd8c2be17-bace-45c1-8eda-d76f16b27a62';
UPDATE questions SET item_order = 5, parent_id = 'd8c2be17-bace-45c1-8eda-d76f16b27a59', research_id = null  WHERE id = '9786d354-32b3-48b7-8e09-8212df7b38df';
UPDATE questions SET item_order = 6, parent_id = 'd8c2be17-bace-45c1-8eda-d76f16b27a59', research_id = null  WHERE id = '29d240d5-295a-499f-a557-5cd6ec6a94e2';

delete from questions where parent_id = 'd8c2be17-bace-45c1-8eda-d76f16b27a47' and id != '04e42ece-8105-4cf7-a2d7-d6b4b75595a4';
UPDATE answer_variants SET show_more_questions = true WHERE id = '5d8de1bc-8495-4bf1-a614-6cb30f965b73';


UPDATE answer_variants SET item_order = 4 WHERE id = '5da683af-5845-4d6b-ad7c-cbb8dc5b9574';


update questions set value_type_id = 'fc00cc5a-f7a5-4974-ad57-9432656d5e0e' where id = 'd8c2be17-bace-45c1-8eda-d76f16b27a62';
