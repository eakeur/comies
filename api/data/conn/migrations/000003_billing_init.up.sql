begin;

create table bills
(
    id           bigint       not null,
    date         timestamp    not null,
    name         varchar(200) not null,
    reference_id bigint       not null,

    primary key (id),
    unique(reference_id)
);

create table bill_items
(
    id           bigint       not null,
    bill_id      bigint       not null,
    reference_id bigint       not null,
    name         varchar(200) not null,
    unit_price   bigint       not null,
    quantity     bigint       not null,
    discounts    bigint       not null,
    

    primary key (id),
    foreign key (bill_id) references bills (id)
);
