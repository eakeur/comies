begin;

create table orders
(
    id                bigint      not null,
    placed_at         timestamp   not null,
    delivery_type     int         not null,
    observations      text        not null,
    customer_name     text        not null,
    customer_phone    varchar(11) not null,
    customer_address  text        not null,

    primary key (id)
);

create table statuses
(
    order_id    bigint    not null,
    value       int       not null,
    occurred_at timestamp not null,

    foreign key (order_id) references orders (id),
    unique (value, order_id)
);

create view latest_statuses as
(
    select 
        distinct on (s.order_id) order_id,
        s.value,
        s.occurred_at
    from statuses s
    order by 
        s.order_id,
        s.occurred_at desc
);

create table items
(
    id           bigint      not null,
    order_id     bigint      not null,
    status       int         not null,
    price        bigint      not null,
    product_id   bigint      not null,
    quantity     bigint      not null,
    observations text        not null,

    primary key (id),
    foreign key (order_id) references orders (id),
    unique (product_id, order_id)
);

end;