begin;

create table products
(
    id               bigint      not null,
    code             varchar(12) not null,
    name             varchar(60) not null,
    type             int         not null,
    cost_price       bigint      not null,
    minimum_sale     bigint      not null,
    sale_unit        varchar(3)  not null,
    maximum_quantity bigint      not null,
    minimum_quantity bigint      not null,
    location         text        not null,

    primary key (id),
    unique (code)
);

create table ingredients
(
    product_id    bigint  not null,
    ingredient_id bigint  not null,
    quantity      bigint  not null,
    optional      boolean not null,

    foreign key (product_id) references products (id),
    foreign key (ingredient_id) references products (id),
    unique (product_id, ingredient_id)
);

create table movements
(
    id         bigint    not null,
    product_id bigint    not null,
    type       int       not null,
    date       timestamp not null,
    quantity   bigint    not null,
    agent_id   bigint    not null,

    primary key (id),
    foreign key (product_id) references products (id)
);

create table prices
(
    id         bigint    not null,
    target_id  bigint    not null,
    value      bigint    not null,
    date       timestamp not null,

    primary key (id)
);
