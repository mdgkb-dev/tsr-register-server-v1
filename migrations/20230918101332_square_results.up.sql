insert into formula_results (name, formula_id, low_range, high_range ,result, color)
-- дюбуа, мостеллер, хэйкок
values
    ('Ниже нормы','79a190cf-b87c-4e23-af8c-8379dc1d50bc', 0, 0.251, 'Ниже нормы', 'red'),
    ('Ниже нормы','79a190cf-b87c-4e23-af8c-8379dc1d50bc', 0.251, 0.251, 'В норме', 'green'),
    ('Ниже нормы','79a190cf-b87c-4e23-af8c-8379dc1d50bc', 0.251, 1000, 'Ниже нормы', 'red'),

    ('Ниже нормы','2f80780d-ab27-4657-bdf9-0110613c050e', 0, 0.247, 'Ниже нормы', 'red'),
    ('Ниже нормы','2f80780d-ab27-4657-bdf9-0110613c050e', 0.247, 0.247, 'В норме', 'green'),
    ('Ниже нормы','2f80780d-ab27-4657-bdf9-0110613c050e', 0.247, 1000, 'Ниже нормы', 'red'),

    ('Ниже нормы','9a7c7609-9463-4215-bab6-4094b20d404b', 0, 0.252, 'Ниже нормы', 'red'),
    ('Ниже нормы','9a7c7609-9463-4215-bab6-4094b20d404b', 0.252, 0.252, 'В норме', 'green'),
    ('Ниже нормы','9a7c7609-9463-4215-bab6-4094b20d404b', 0.252, 1000, 'Ниже нормы', 'red');



UPDATE public.formulas SET formula = '(0.024265*(weight^0.5378))*(height^0.3964)' WHERE id = '9a7c7609-9463-4215-bab6-4094b20d404b';
UPDATE public.formulas SET formula = '(0.0167*(weight^0.5))*(height^0.5)' WHERE id = '2f80780d-ab27-4657-bdf9-0110613c050e';
UPDATE public.formulas SET formula = '(0.007184*(weight^0.425))*(height^0.725)' WHERE id = '79a190cf-b87c-4e23-af8c-8379dc1d50bc';
