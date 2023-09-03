insert into public.researches_domains (id, research_id, domain_id)
values  ('c8db37cc-cf86-4574-8c2e-e0ac3251cc86', 'c21eb5de-e180-4faa-b1eb-bfd77dddb470', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');

alter table researches
add column is_laboratory bool default true;


UPDATE public.researches SET is_laboratory = false WHERE id = 'e1ac2248-1962-4301-9a2a-34a299e26934';
UPDATE public.researches SET is_laboratory = false WHERE id = 'c21eb5de-e180-4faa-b1eb-bfd77dddb470';
UPDATE public.researches SET is_laboratory = false WHERE id = 'e9f2300f-afb7-43e0-93b9-eb110edfa689';
UPDATE public.researches SET is_laboratory = false WHERE id = 'dc000421-4a9b-40f4-a6ad-a18b08e5cd6e';
UPDATE public.researches SET is_laboratory = false WHERE id = '8f2f58fd-38e2-4644-b8e7-e05794e838a8';
UPDATE public.researches SET is_laboratory = false WHERE id = 'd9fa9763-fd9a-4123-a55e-790774894e98';
UPDATE public.researches SET is_laboratory = false WHERE id = 'e9f2300f-afb7-43e0-93b9-eb110edfa685';
UPDATE public.researches SET is_laboratory = false WHERE id = 'e9f2300f-afb7-43e0-93b9-eb110edfa686';
UPDATE public.researches SET is_laboratory = false WHERE id = 'f14494e4-e048-4bf7-9894-3955a11368df';
UPDATE public.researches SET is_laboratory = false WHERE id = 'e9f2300f-afb7-43e0-93b9-eb110edfa683';

insert into public.researches_domains (id, research_id, domain_id)
values  ('88a8790c-89df-447a-bd2a-4002f31694ea', '8f2f58fd-38e2-4644-b8e7-e05794e838a8', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('c8860bad-27e5-421a-adca-bcc987a00482', 'e9f2300f-afb7-43e0-93b9-eb110edfa683', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('bad9e9cf-bf24-4e17-83e6-a10d083e653f', 'e9f2300f-afb7-43e0-93b9-eb110edfa686', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('12efe5e1-9847-4da4-a298-114e3d381ecf', 'e9f2300f-afb7-43e0-93b9-eb110edfa689', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');