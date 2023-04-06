create table realm
(
    id   uuid  not null default gen_random_uuid()
        primary key,
    data jsonb not null default '{}'
);

create table app
(
    id       uuid    not null default gen_random_uuid()
        primary key,
    realm_id uuid    not null
        constraint app_fk_realm_id
            references realm (id) on update cascade on delete cascade,
    active   boolean not null default true,
    data     jsonb   not null default '{}'
);
create index app_idx_realm_id on app (realm_id);
create index app_idx_active on app (active);

create table "role"
(
    id         uuid    not null          default gen_random_uuid()
        primary key,
    app_id     uuid
        constraint role_fk_app_id
            references app (id) on update cascade on delete cascade,
    is_fetched boolean not null          default false,
    code       text    not null not null default '',
    dsc        text    not null          default ''
);
create index role_idx_app_id on "role" (app_id);
create index role_idx_is_fetched on "role" (is_fetched);

create table endpoint
(
    id     uuid    not null default gen_random_uuid()
        primary key,
    app_id uuid    not null
        constraint endpoint_fk_app_id
            references app (id) on update cascade on delete cascade,
    active boolean not null default true,
    data   jsonb   not null default '{}'
);
create index endpoint_idx_app_id on endpoint (app_id);
create index endpoint_idx_active on endpoint (active);

-------------------------------------------------------------------------------

insert into realm(id, data)
select id, val -> 'general' || jsonb_build_object('name', name)
from data;

insert into app(id, realm_id, active, data)
select (a ->> 'id')::uuid, d.id, (a -> 'active')::bool, (a - 'id' - 'endpoints')
from data d,
     jsonb_array_elements(d.val -> 'apps') as a;

insert into endpoint(id, app_id, active, data)
select (e ->> 'id')::uuid, (a ->> 'id')::uuid, (e -> 'active')::bool, e - 'id'
from data d,
     jsonb_array_elements(d.val -> 'apps') as a,
     jsonb_array_elements(a -> 'endpoints') as e;
