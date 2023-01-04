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

create table data
(
    id   uuid  not null default gen_random_uuid()
        primary key,
    name text  not null default '',
    val  jsonb not null default '{}'
);

do
$$
    declare
    begin
    end ;
$$;
