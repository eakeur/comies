begin;

    create table member (
        id                      bigint                   not null,
        active                  boolean                  not null,
        name                    varchar(24)              not null,
        nickname                varchar(24)              not null,
        password                text                     not null,

        constraint member_pk primary key (id),
        constraint nickname_uk unique (nickname)
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

        constraint products_pk primary key (id),
        constraint code_store_id_uk unique (code)
    );

    create table ingredients (
        id bigint not null,
        product_id bigint not null,
        ingredient_id bigint not null,
        quantity bigint not null,
        optional boolean not null,

        constraint ingredients_pk primary key (id),
        constraint product_id_fk foreign key (product_id) references products(id) on delete cascade,
        constraint ingredient_id_fk foreign key (ingredient_id) references products(id),
        constraint product_ingredient_id_uk unique (product_id, ingredient_id)
    );

    create table orders (
        id bigint not null,
        identification text not null,
        placed_at timestamp with time zone not null,
        delivery_mode varchar(30) not null,
        observations text,
        address text not null,
        phone varchar(11) not null,

        constraint orders_pk primary key (id)
    );

    create table orders_flow (
        id bigint not null,
        order_id bigint not null,
        status varchar(30) not null,
        occurred_at timestamp with time zone not null,

        constraint flow_pk primary key (id),
        constraint order_id_fk foreign key (order_id) references orders(id),
        constraint order_status_uk unique(status, order_id)
    );

    create view orders_statuses as (
         select distinct on (f.order_id)
            f.order_id,
            f.status
        from
            orders_flow f
        order by
            f.order_id,
            f.id desc,
            f.occurred_at desc nulls last
    );

    create table items (
        id bigint not null,
        order_id bigint not null,
        status varchar(30) not null,
        price bigint not null,
        product_id bigint not null,
        quantity bigint not null,
        observations text not null,

        constraint items_pk primary key (id),
        constraint order_id_fk foreign key (order_id) references orders(id),
        constraint product_order_uk unique(product_id, order_id)
    );

    create table item_details (
        item_id bigint not null,
        ignored bigint,
        replacing bigint,
        replaced bigint,

        constraint item_id_fk foreign key (item_id) references items(id)
    );

    create table stocks (
        id bigint not null,
        active boolean not null,
        target_id bigint not null,
        maximum_quantity bigint not null,
        minimum_quantity bigint not null,
        location text,

        constraint stocks_id primary key (id),
        constraint target_store_uk unique (target_id)
    );

    create table movements (
        id bigint not null,
        stock_id bigint not null,
        type varchar(30) not null,
        date timestamp with time zone not null,
        quantity bigint not null,
        value bigint not null,
        agent_id bigint not null,

        constraint movements_id primary key (id),
        constraint stock_id_fk foreign key (stock_id) references stocks(id)
    );





end;