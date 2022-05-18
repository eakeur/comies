begin;

    create table stores (
        id        bigint       not null,
        active    boolean      not null,
        name      text         not null,
        nickname  varchar(24)  not null,
        document  text         not null,

        constraint stores_pk primary key id,
        constraint nickname_uk unique nickname
    );

    create table crew (
        id                      bigint                   not null,
        active                  boolean                  not null,
        name                    varchar(24)              not null,
        full_name               text                     not null,
        nickname                varchar(24)              not null,
        picture                 text                     not null,
        password                text                     not null,
        password_changed_at     timestamp with time zone not null,
        last_seen               timestamp with time zone not null,
        store_id                bigint                   not null,

        constraint crew_pk primary key id,
        constraint stores_store_id foreign key store_id,
        constraint nickname_store_id_uk unique (nickname, store_id)
    );

    create table products (
        id           bigint      not null,
        active       boolean     not null,
        code         varchar(12) not null,
        name         varchar(60) not null,
        type         int         not null,
        cost_price   bigint      not null,
        sale_price   bigint      not null,
        minimum_sale bigint      not null,
        sale_unit    varchar(3)  not null,
        store_id     bigint      not null,

        constraint products_pk primary key id,
        constraint code_store_id_uk unique (code, store_id)
    );




end;