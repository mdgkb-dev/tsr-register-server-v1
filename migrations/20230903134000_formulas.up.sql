alter table formulas add column xlsx bool default true;

insert into public.formulas (id, name, formula, research_id, age_relation, sex_relation, color, xlsx)
values  ('79a190cf-b87c-4e23-af8c-8379dc1d50bc', 'Площадь тела Дюбуа и Дюбуа', '(0.007184*weight^0.425)*(height^0.725)', 'af81a162-90fa-40f2-801c-f8eea0bfa947', true, true, '#323d70', true),
        ('c5fe9e54-4ba4-4306-9485-37271623d535', 'Вес', 'weight', 'af81a162-90fa-40f2-801c-f8eea0bfa947', false, false, '#a8b536', false),
        ('68068485-f927-41b5-ab43-289a6a3978f5', 'Рост', 'height', 'af81a162-90fa-40f2-801c-f8eea0bfa947', false, false, '#1c5920', false),
        ('2f80780d-ab27-4657-bdf9-0110613c050e', 'Площадь тела Мостеллера', '((0,0167*weight)^0,5)*(height^0.5)', 'af81a162-90fa-40f2-801c-f8eea0bfa947', true, true, '#323d70', true),
        ('9a7c7609-9463-4215-bab6-4094b20d404b', 'Площадь тела Хейкока', '((0.024265*weight)^0.5378)*(height^0.3964)', 'af81a162-90fa-40f2-801c-f8eea0bfa947', true, true, '#323d70', true);

update formulas set xlsx = false where name in ('Вес', 'Рост');

