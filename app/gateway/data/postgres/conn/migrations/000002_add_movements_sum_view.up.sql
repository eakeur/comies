begin;
create view products_balances as
select
    product_id,
    sum(
            case
                when type = 10
                    then quantity
                else  -1 * quantity
                end
        )
        as balance
from movements
group by product_id;

end;
