alter table public.cart
    add constraint pk_cart_id primary key (id),
    alter column price set not null,
    add check (price >= 0);

alter table public.user
    add constraint pk_user_id primary key (id),
    add unique (cart_id),
    add constraint fk_user_cart_id
    foreign key (cart_id)
    references public.cart(id) on delete cascade,
    alter column name set not null,
    alter column surname set not null,
    alter column email set not null,
    add unique (email),
    alter column password set not null,
    alter column role set not null;

alter table public.product
    add constraint pk_product_id primary key (id),
    alter column name set not null,
    alter column description set not null,
    alter column price set not null,
    add check (price > 0),
    alter column category set not null;

create index idx_product_name on public.product (name);

alter table public.cart_product
    add constraint pk_cart_product_id primary key (id),
    alter column cart_id set not null,
    alter column product_id set not null,
    add constraint fk_cart_product_cart_id
    foreign key (cart_id)
    references public.cart(id) on delete cascade,
    add constraint fk_cart_product_product_id
    foreign key (product_id)
    references public.product(id) on delete cascade,
    add constraint uc_cart_product unique (cart_id,product_id),
    alter column quantity set not null,
    add check (quantity > 0);

alter table public.shop
    add constraint pk_shop_id primary key (id),
    alter column seller_id set not null,
    alter column name set not null,
    alter column description set not null,
    alter column requisites set not null,
    alter column email set not null,
    add unique (email),
    add constraint fk_shop_seller_id
    foreign key (seller_id)
    references public.user(id) on delete cascade;

alter table public.shop_product
    add constraint pk_shop_product_id primary key (id),
    alter column shop_id set not null,
    add constraint fk_shop_product_shop_id
    foreign key (shop_id)
    references public.shop(id) on delete cascade,
    alter column product_id set not null,
    add constraint fk_shop_product_product_id
    foreign key (product_id)
    references public.product(id) on delete cascade,
    alter column quantity set not null,
    add check (quantity >= 0),
    add constraint uc_shop_product unique (shop_id,product_id);

alter table public.withdraw
    add constraint pk_withdraw_id primary key (id),
    alter column shop_id set not null,
    add constraint fk_withdraw_shop_id
    foreign key (shop_id)
    references public.shop(id) on delete cascade,
    alter column comment set not null,
    alter column sum set not null,
    add check (sum > 0),
    alter column status set not null;

alter table public.order_customer
    add constraint pk_order_customer_id primary key (id),
    alter column customer_id set not null,
    add constraint fk_order_customer_customer_id
    foreign key (customer_id)
    references public.user(id) on delete cascade,
    alter column address set not null,
    alter column created_at set not null,
    alter column total_price set not null,
    add check (total_price > 0),
    alter column payed set not null;

alter table public.order_shop
    add constraint pk_order_shop_id primary key (id),
    alter column shop_id set not null,
    add constraint fk_order_shop_shop_id
    foreign key (shop_id)
    references public.shop(id) on delete cascade,
    alter column order_customer_id set not null,
    add constraint fk_order_shop_order_customer_id
    foreign key (order_customer_id)
    references public.order_customer(id) on delete cascade,
    alter column status set not null,
    alter column notified set not null,
    add constraint uc_order_shop unique (shop_id,order_customer_id);

alter table public.order_shop_product
    add constraint pk_order_shop_product_id primary key (id),
    alter column order_shop_id set not null,
    add constraint fk_order_shop_product_order_shop_id
    foreign key (order_shop_id)
    references public.order_shop(id) on delete cascade,
    alter column product_id set not null,
    add constraint fk_order_shop_product_product_id
    foreign key (product_id)
    references public.product(id) on delete cascade,
    alter column quantity set not null,
    add check (quantity >= 0),
    add constraint uc_order_shop_product unique (order_shop_id,product_id);
