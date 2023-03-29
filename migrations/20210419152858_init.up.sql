do
$$
    begin
        execute 'ALTER DATABASE ' || current_database() || ' SET timezone = ''+06''';
    end;
$$;

create table cfg
(
    v jsonb not null default '{}'
);

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

do
$$
    declare
    begin
    end ;
$$;
