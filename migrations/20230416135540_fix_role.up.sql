truncate table "role" restart identity cascade;

alter table "role"
    add column realm_id uuid not null
        constraint role_fk_realm_id
            references realm (id) on update cascade on delete cascade;
