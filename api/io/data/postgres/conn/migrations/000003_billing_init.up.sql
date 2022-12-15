begin;

create table bills
(
    id           bigint       not null,
    date         timestamp    not null,
    name         varchar(200) not null,
    reference_id bigint       not null,

    primary key (id)
);

create table bill_items
(
    id           bigint    not null,
    bill_id      bigint    not null
    date         timestamp not null,
    description  text      not null,
    credits      bigint    not null,
    debts        bigint    not null
    reference_id bigint    not null,

    primary key (id),
    foreign key (bill_id) references bills (id)
);
