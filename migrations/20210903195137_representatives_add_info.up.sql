alter table representative
    add created_at timestamp default current_timestamp not null;

alter table representative
    add created_by_id uuid;

alter table representative
    add updated_at timestamp default current_timestamp;

alter table representative
    add updated_by_id uuid;

alter table representative
    add constraint representatives_users_id_fk
        foreign key (created_by_id) references users
            on update restrict on delete restrict;

alter table representative
    add constraint representatives_users_id_fk_2
        foreign key (updated_by_id) references users
            on update restrict on delete restrict;

