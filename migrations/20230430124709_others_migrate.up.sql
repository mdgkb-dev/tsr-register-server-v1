insert into questions (id, name, item_order, value_type_id, parent_id)
select answer_comments.id, answer_comments.name, answer_comments.item_order,
       case
--     when answer_comments.name in ('Указать год постановки диагноза', 'С какого возраста') then 'efdd456c-091b-49d9-ac32-d0d345f88e64'::uuid
           when answer_comments.name in ('На сколько дней', 'Указать количество эпизодов') then '47affcc5-5d32-4b1f-bf07-33382ed06cda'::uuid
           else '9f61f302-6821-40b9-94bc-78dedf955a11'::uuid
           end,
       av.question_id
from answer_comments
         join answer_variants av on answer_comments.answer_variant_id = av.id
         join questions q on q.id = av.question_id;


insert into answers (id, value_string,  question_id, research_result_id)
select r.id,r.value,r.register_property_other_id,register_group_to_patient_id
from register_property_other_to_patient r
         join questions q on r.register_property_other_id = q.id
         join value_type vt on vt.id = q.value_type_id and vt.name = 'string';


insert into answers (id, value_number,  question_id, research_result_id)
select r.id,regexp_replace(r.value, '[^0-9]+', '', 'g')::numeric,r.register_property_other_id,register_group_to_patient_id
from register_property_other_to_patient r
         join questions q on r.register_property_other_id = q.id
         join value_type vt on vt.id = q.value_type_id and vt.name = 'number'
where regexp_replace(r.value, '[^0-9]+', '', 'g') != '';


update answer_variants
set show_more_questions = true
where id in
      (select answer_variants.id from answer_variants
                                          join questions q on answer_variants.question_id = q.id
                                          join questions q2 on q.id = q2.parent_id
       where answer_variants.name not ilike ('%нет%')
         and answer_variants.name not ilike ('%неизвестна%')
         and answer_variants.name not ilike ('%мин.%')
         and answer_variants.name not ilike ('%минут%')
         and answer_variants.name not ilike ('%1 час%')
         and answer_variants.name not ilike ('%Физическая нагрузка%')
         and answer_variants.name not ilike ('%Школа%')
         and answer_variants.name not ilike ('%Дом%')
         and answer_variants.name not ilike ('%Работа%')
         and answer_variants.name not ilike ('%Общественно%')
         and answer_variants.name not ilike ('%в гостях%')
         and answer_variants.name not ilike ('%учреждение%')
         and answer_variants.name not ilike ('%физическая нагр%')
         and answer_variants.name not ilike ('%темперар%')
         and answer_variants.name not ilike ('%алкоголь%')
         and answer_variants.name not ilike ('%стресс%')
         and answer_variants.name not ilike ('%меструа%')
         and answer_variants.name not ilike ('%орви%')
         and answer_variants.name not ilike ('%вдыхание%')
         and answer_variants.name not ilike ('%инъекционно%')
         and answer_variants.name not ilike ('%контакт%')
         and answer_variants.name not ilike ('%употребление%')
         and answer_variants.name not ilike ('%запах%')
         and answer_variants.name not ilike ('%сыпь%')
         and answer_variants.name not ilike ('%жара%')
         and answer_variants.name not ilike ('%покраснение%')
         and answer_variants.name not ilike ('%отёк%')
         and answer_variants.name not ilike ('%осиплость%')
         and answer_variants.name not ilike ('%затруднение%')
         and answer_variants.name not ilike ('%исчезновение%')
         and answer_variants.name not ilike ('%кашель%')
         and answer_variants.name not ilike ('%свисты%')
         and answer_variants.name not ilike ('%удушье%')
         and answer_variants.name not ilike ('%головокружение%')
         and answer_variants.name not ilike ('%потеря%')
         and answer_variants.name not ilike ('%снижение%')
         and answer_variants.name not ilike ('%рвота%')
         and answer_variants.name not ilike ('%боли в животе%')
         and answer_variants.name not ilike ('%диарея%')
         and answer_variants.name not ilike ('%сонливость%')
         and answer_variants.name not ilike ('%острая потеря%')
         and answer_variants.name not ilike ('%спутанность%')
         and answer_variants.name not ilike ('%адреналин%')
         and answer_variants.name not ilike ('% мин%')
         and answer_variants.name not ilike ('%анафилаксия%')
         and answer_variants.name not ilike ('%крапивница%')
         and answer_variants.name not ilike ('%ангиоотёк%')
         and answer_variants.name not ilike ('%проконсультиро%')
         and answer_variants.name not ilike ('%исключать%')
         and answer_variants.name not ilike ('%пройти%')
         and answer_variants.name not ilike ('%носить%')
         and answer_variants.name not ilike ('%кожные%')
         and answer_variants.name not ilike ('%уровень%')
         and answer_variants.name not ilike ('%крапивница%')
         and answer_variants.name not ilike ('%анафилаксия%'));


update answers set question_id = '754bcf9d-fb22-46fa-a00e-c6314881d0a8' where question_id in ('cc330c04-ea13-4c13-b09c-35b2315426f4','ca713203-aed7-4f6d-a996-d4942d9fe7e9');
update answers set question_id = '14cf7d8a-ddf4-48df-ab6d-388efbb28807' where question_id in ('ee09327f-3627-42a1-8376-a84cb15445bf','4ca057c8-ad49-455b-b3e3-93801d03c378');
update answers set question_id = '8c361dee-867d-44e4-a2a8-4a59aa1a1183' where question_id in ('93ed4ed7-3316-40fd-b342-131ff1432fe9','ae8ecc65-dd26-4b0d-8de1-0e83a69fff6d');
update answers set question_id = 'fbf19484-5d53-4a97-86d5-979e66846f58' where question_id in ('5d6a6566-3d91-4327-8eb6-52cb9d966d17','a70ff1f2-b3a3-4987-8d17-553066767d6e');
update answers set question_id = 'aac5f077-c37d-4dbc-bcf0-22f93e41bf68' where question_id in ('fba80068-895a-47cb-aa79-b5c7df3de80a','7396b8a5-d5a5-4212-bf21-4d2f4fd32814');
update answers set question_id = 'e723405c-9198-43a8-b014-8248405444f0' where question_id in ('7b98630b-0855-49a6-8566-bf9f3e440ddf','79d18dc0-56e0-4ae3-8756-2a2e9c25f040');
update answers set question_id = '444ca832-c844-4840-be23-7cce43a6e9a8' where question_id in ('57cdb3e4-473e-4018-853c-ca631be6b486','9a6af763-1c12-44e5-8127-9ab0ba2d369d');
delete from questions where id in (
                                   'cc330c04-ea13-4c13-b09c-35b2315426f4','ca713203-aed7-4f6d-a996-d4942d9fe7e9',
                                   'ee09327f-3627-42a1-8376-a84cb15445bf','4ca057c8-ad49-455b-b3e3-93801d03c378',
                                   '93ed4ed7-3316-40fd-b342-131ff1432fe9','ae8ecc65-dd26-4b0d-8de1-0e83a69fff6d',
                                   '5d6a6566-3d91-4327-8eb6-52cb9d966d17','a70ff1f2-b3a3-4987-8d17-553066767d6e',
                                   'fba80068-895a-47cb-aa79-b5c7df3de80a','7396b8a5-d5a5-4212-bf21-4d2f4fd32814',
                                   '7b98630b-0855-49a6-8566-bf9f3e440ddf','79d18dc0-56e0-4ae3-8756-2a2e9c25f040',
                                   '57cdb3e4-473e-4018-853c-ca631be6b486','9a6af763-1c12-44e5-8127-9ab0ba2d369d'
    );


update answers set question_id = 'd2dfe3e5-622f-4e40-9191-c9db18de0fe3' where question_id in ('334d8f74-7c1c-4555-b7e6-693b6b3b04c4','4f777984-278e-47b3-ae99-859047533f10');
update answers set question_id = 'df29c5d1-296a-4510-991a-21ac33c1b87e' where question_id in ('a549a70d-2d5e-4073-a3b6-0b1da4ae4e06','9b7176ef-1f8b-425d-b358-75a0c15af14d');
update answers set question_id = 'e8b60848-6291-4406-8eca-f8ad8c304e87' where question_id in ('d1fe355b-ef4a-4ec9-9509-1c150b315f5d','ce9fab45-9faa-4f01-9e32-f2c56d8e82a6');
update answers set question_id = '5f6a68a3-1850-4b9a-a304-0466256386dc' where question_id in ('6df68f5c-820d-4151-be89-9fc467836d25','42252a39-7875-49fe-bc01-4eda4d31d1b8');
update answers set question_id = '553f9879-2809-4631-a638-7ff9410f251d' where question_id in ('322c58b2-25e3-4104-bdce-8ae294d2aaa3','19b461c8-1c01-4a84-a73b-17b83dad0803');
update answers set question_id = 'd5a07388-c214-4529-b734-1b1346beb33f' where question_id in ('83814c29-421c-450b-b9be-e5d07708bea2','9f1bfbe2-54cd-4b84-9591-581b72939a4d');
update answers set question_id = '794b1654-7070-442e-802a-eb2c9b378ee5' where question_id in ('9cdb4a90-9f3f-4d47-ba5a-92ec45bb82f6','de95f6dd-493b-4f03-9f08-3ac6d38ac439');
delete from questions where id in (
                                   '334d8f74-7c1c-4555-b7e6-693b6b3b04c4','4f777984-278e-47b3-ae99-859047533f10',
                                   'a549a70d-2d5e-4073-a3b6-0b1da4ae4e06','9b7176ef-1f8b-425d-b358-75a0c15af14d',
                                   'd1fe355b-ef4a-4ec9-9509-1c150b315f5d','ce9fab45-9faa-4f01-9e32-f2c56d8e82a6',
                                   '6df68f5c-820d-4151-be89-9fc467836d25','42252a39-7875-49fe-bc01-4eda4d31d1b8',
                                   '322c58b2-25e3-4104-bdce-8ae294d2aaa3','19b461c8-1c01-4a84-a73b-17b83dad0803',
                                   '83814c29-421c-450b-b9be-e5d07708bea2','9f1bfbe2-54cd-4b84-9591-581b72939a4d',
                                   '9cdb4a90-9f3f-4d47-ba5a-92ec45bb82f6','de95f6dd-493b-4f03-9f08-3ac6d38ac439'
    );


update answers set question_id = '31035f17-7320-49b0-a26b-72e3aedd425e' where question_id in (
                                                                                              '481f11e0-e8cb-43e7-8ff6-1ca4aa3031cd','09c8701b-76be-4008-876e-fe536f2176ae', '18f11adc-313b-41dc-8232-cdaef6af0599', '032757d6-d0d2-4e53-aee5-f59a68851372',
                                                                                              '1103b6e1-ead8-4c71-8eb5-b37327417e6d');
update answers set question_id = '1eb310cf-b0bc-4410-9a69-6f50dbdb8b1d' where question_id in ('c4e58c07-a802-4c7a-a916-b2bbe1f4d646');
delete from questions where id in (
                                   '481f11e0-e8cb-43e7-8ff6-1ca4aa3031cd','09c8701b-76be-4008-876e-fe536f2176ae', '18f11adc-313b-41dc-8232-cdaef6af0599', '032757d6-d0d2-4e53-aee5-f59a68851372',
                                   '1103b6e1-ead8-4c71-8eb5-b37327417e6d','c4e58c07-a802-4c7a-a916-b2bbe1f4d646'
    );



update answers set question_id = 'd9f933ad-b11e-4c82-b523-ac2f46361ed1' where question_id in (
                                                                                              '94227bab-3197-49f5-b327-4a2ad3129603','62beb7b0-4cbc-4bd2-adeb-65dbeabf9867','6a63c510-4c39-49e3-85dd-a77586fa629a','ec80417f-fad8-4010-9fa3-01857e0bb0d9',
                                                                                              '8797245b-35ff-458a-a10f-1a1d28ca27dc','32f70724-766c-47b4-8ba4-3fb7b4d69743'                                                 );
update answers set question_id = '04e42ece-8105-4cf7-a2d7-d6b4b75595a4' where question_id in ('0ab865cc-cb3c-4339-8fa5-3a42610a6b52','81f831e2-3789-4489-ae91-6ccbd4b3ed5d','8f29291a-0ae1-4280-9890-ef3d81955cd1');
-- update answers set question_id = '' where question_id in ('',);
delete from questions where id in ( '94227bab-3197-49f5-b327-4a2ad3129603','62beb7b0-4cbc-4bd2-adeb-65dbeabf9867','6a63c510-4c39-49e3-85dd-a77586fa629a','ec80417f-fad8-4010-9fa3-01857e0bb0d9',
                                    '8797245b-35ff-458a-a10f-1a1d28ca27dc','32f70724-766c-47b4-8ba4-3fb7b4d69743','0ab865cc-cb3c-4339-8fa5-3a42610a6b52','81f831e2-3789-4489-ae91-6ccbd4b3ed5d','8f29291a-0ae1-4280-9890-ef3d81955cd1'      );