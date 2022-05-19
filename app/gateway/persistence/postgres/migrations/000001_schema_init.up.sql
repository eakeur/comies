begin;

    create table stores (
        id        bigint       not null,
        active    boolean      not null,
        name      text         not null,
        nickname  varchar(24)  not null,

        constraint stores_pk primary key id,
        constraint nickname_uk unique nickname
    );

    create table crew (
        id                      bigint                   not null,
        active                  boolean                  not null,
        name                    varchar(24)              not null,
        nickname                varchar(24)              not null,
        password                text                     not null,
        store_id                bigint                   not null,

        constraint crew_pk primary key id,
        constraint stores_store_id foreign key store_id references store(id),
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

    create table ingredients (
        id bigint not null,
        active boolean not null,
        product_id bigint not null,
        ingredient_id bigint not null,
        quantity bigint not null,
        optional boolean not null,
        store_id bigint not null,

        constraint ingredients_pk primary key id,
        constraint product_id_fk foreign key product_id references products(id),
        constraint ingredient_id_fk foreign key ingredient_id references products(id),
        constraint product_ingredient_id_uk unique (product_id, ingredient_id, store_id)
    );

    create table orders (
        id bigint not null,
        active boolean not null,
        identification text not null,
        placed_at timestamp with time zone not null,
        status int not null,
        delivery_mode int not null,
        observation text,
        store_id bigint not null,

        constraint orders_pk primary key id
    );

    create table items (
        id bigint not null,
        active bigint not null,
        order_id bigint not null,
        status int not null,
        price bigint not null,
        product_id bigint not null,
        quantity bigint not null,
        observations text not null,
        store_id bigint not null,

        constraint items_pk primary key id,
        constraint order_id_kf foreign key order_id references orders(id),
        constraint product_order_uk unique(product_id, order_id)
    );

    create table item_details (
        item_id bigint not null,
        ignored bigint,
        replacing bigint,
        replaced bigint,

        constraint item_id_fk foreign key item_id references items(id)
    );

    create table stocks (
        id bigint not null,
        active boolean not null,
        target_id bigint not null,
        maximum_quantity bigint not null,
        minimum_quantity bigint not null,
        location text,
        store_id bigint not null,

        constraint stocks_id primary key id,
        constraint target_store_uk unique (target_id, store_id)
    );

    create table movements (
        id bigint not null,
        active boolean not null,
        stock_id bigint not null,
        type int not null,
        date timestamp with time zone not null,
        quantity bigint not null,
        value bigint not null,
        agent bigint not null,

        constraint movements_id primary key id,
        constraint stock_id_fk foreign key stock_id references stock(id)
    );





end;