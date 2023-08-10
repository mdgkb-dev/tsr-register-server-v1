INSERT INTO public.researches_pools (id, name) VALUES ('d9e8d64c-7587-4d7e-b8b9-5bf5cff369eb', 'Дополнительные сведения');

INSERT INTO public.researches (id, name, with_dates, with_scores) VALUES ('f14494e4-e048-4bf7-9894-3955a11368df', 'Дополнительные сведения', false, false);

update questions q set research_id = 'f14494e4-e048-4bf7-9894-3955a11368df' where q.id = 'b54b5ce1-998b-4363-8ddb-e5dddba237bb';
update questions q set research_id = 'f14494e4-e048-4bf7-9894-3955a11368df' where q.id = 'b54b5ce1-998b-4363-8ddb-e5dddba237bc';
update questions q set item_order = 1 where q.id = 'b54b5ce1-998b-4363-8ddb-e5dddba237bc';

INSERT INTO public.researches_pools_researches (id, researches_pool_id, research_id, item_order) VALUES ('05e0c6bd-8d1c-4356-9877-cb2e85f16bb1', 'd9e8d64c-7587-4d7e-b8b9-5bf5cff369eb', 'f14494e4-e048-4bf7-9894-3955a11368df', 0);

update questions q set domain_id = null where q.id = 'b54b5ce1-998b-4363-8ddb-e5dddba237bc';