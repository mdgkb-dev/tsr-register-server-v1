create table custom_sections(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    component varchar,
    item_order integer
);

create table custom_sections_domains
(
    id uuid default uuid_generate_v4() not null primary key,
    custom_section_id uuid not null references custom_sections,
    domain_id uuid not null references domains
);


create table menus_domains
(
    id uuid default uuid_generate_v4() not null primary key,
    menu_id uuid not null references menus,
    domain_id uuid not null references domains
);



insert into public.custom_sections (id, name, component, item_order)
values  ('3425c483-7bef-4d1a-9899-10a754b72fff', 'info', 'Паспортные данные', 0),
        ('e44eccfa-f3ec-4a32-b1f5-f643ad7baab9', 'diagnosis', 'Диагнозы', 1),
        ('c0fda316-3896-46ee-a9e2-25b6262d9314', 'anamneses', 'Анамнез', 2),
        ('b232baa0-ad7b-4245-8f6d-9ec034f7df58', 'patientResearches', 'Исследования', 3),
        ('77fe2abb-2ced-4afb-87b2-6bd2d27b87b4', 'representatives', 'Представители', 4),
        ('898a14c5-1f07-4811-8f32-5bbe5a4252b9', 'disability', 'Инвалидность', 5),
        ('112be431-4a13-4ced-8624-1cce1bd04e78', 'documents', 'Документы', 6),
        ('40b1bd9f-6a6b-401e-9a59-40504a1d5b47', 'commissions', 'Врачебные комиссии', 7),
        ('16994783-844e-4cfa-a296-aba1423db014', 'registers', 'Регистры', 8);


insert into public.custom_sections_domains (id, custom_section_id, domain_id)
values  ('55a3a692-500d-465e-b10c-5539390facc7', '3425c483-7bef-4d1a-9899-10a754b72fff', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('5815ff58-1b92-4195-9d96-94a2f754b333', 'e44eccfa-f3ec-4a32-b1f5-f643ad7baab9', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('d940f883-7c76-4fa3-b59d-9a3c369fb00e', 'c0fda316-3896-46ee-a9e2-25b6262d9314', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('50742f8f-b21f-4b85-9422-1cd83523942a', 'b232baa0-ad7b-4245-8f6d-9ec034f7df58', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('96c2b0fa-fee3-441f-a23f-fa1d213d5f71', '77fe2abb-2ced-4afb-87b2-6bd2d27b87b4', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('3dbd857a-7b2c-4850-8946-e0f97b7c6b56', '898a14c5-1f07-4811-8f32-5bbe5a4252b9', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('838dd140-a40a-4f98-b366-5048fd408659', '112be431-4a13-4ced-8624-1cce1bd04e78', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('d58c8258-5fa9-48c2-b161-be7a0afcb5fe', '40b1bd9f-6a6b-401e-9a59-40504a1d5b47', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('f6de52e8-eb1b-4efb-b058-7dc3c14b615b', '16994783-844e-4cfa-a296-aba1423db014', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('3a87e814-bce8-45de-869e-190356afb89b', '3425c483-7bef-4d1a-9899-10a754b72fff', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('e4e2914d-b20c-411e-baf7-9f606ef41493', 'e44eccfa-f3ec-4a32-b1f5-f643ad7baab9', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('a24f1307-e9c0-40b9-bf13-f68954366f0c', 'c0fda316-3896-46ee-a9e2-25b6262d9314', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('cd84b551-8694-4c72-97bb-7cedb19972be', 'b232baa0-ad7b-4245-8f6d-9ec034f7df58', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('3475ceac-170d-4b6b-b098-e1f639f3fac5', '77fe2abb-2ced-4afb-87b2-6bd2d27b87b4', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');


insert into public.menus_domains (id, menu_id, domain_id)
values  ('a15b29ec-3bb0-4d6c-bcb0-a4cd9bbe3894', '408eae9d-043d-4fed-ae88-51f27c508ad1', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('4277072b-edde-4b66-b6e3-972010409514', '343a7237-f28b-4f16-99e2-4bdcd6b78f8b', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('b2f7d7ba-db01-47eb-93d6-977008910bb8', '94acb304-6651-4b6e-9646-11aeec3910a5', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('770659bb-534d-48de-9eb7-c890da126b5a', 'e092080c-15c2-43ba-907e-bfd33e08c8eb', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('b0030cc8-7138-4929-8ad2-b23b41a67ed2', 'ab0fa211-ba61-40d9-a48b-5724cfaa75ee', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('e4bd660e-3f3a-4b63-884c-014c6d9740ae', 'd24a2240-684b-476b-96f9-ca5363f5b384', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('a5deb733-f2fd-41fc-a482-96976e401c87', '408eae9d-043d-4fed-ae88-51f27c508ad1', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('96d9f363-870c-46e8-833f-c33517d3b52f', '343a7237-f28b-4f16-99e2-4bdcd6b78f8b', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');

        UPDATE public.users_accounts SET password = '$2a$10$r3AlsAUhJZASy028tu6GcuitOp6vkL9xKszP1FV/RQ2ZdNagKHKW6' WHERE id = '5ea58a1c-0671-4df6-b46b-72242641fbbb'