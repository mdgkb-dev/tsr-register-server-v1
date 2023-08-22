create table users_accounts
(
    id        uuid default uuid_generate_v4() not null primary key,
    password  varchar not null,
    email     varchar not null unique ,
    uuid      uuid default uuid_generate_v4() not null
);

alter table users
    add column user_account_id uuid REFERENCES users_accounts(id) ON UPDATE CASCADE ON DELETE CASCADE;



with ua as (
    insert into users_accounts (id, password, email, uuid)
        select uuid_generate_v4(), password, email, uuid from users
           returning id, email
)
update users u set user_account_id = ua.id
             from ua where u.email = ua.email;


alter table users drop column email, drop column password, drop column uuid;