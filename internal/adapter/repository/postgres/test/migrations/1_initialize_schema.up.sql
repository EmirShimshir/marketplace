create table public.cart (
    id uuid primary key,
    price bigint not null
);

create type user_role as enum ('Customer', 'Seller', 'Moderator');
create table public.user (
    id uuid primary key,
    cart_id uuid unique not null,
    name varchar(255) not null,
    surname varchar(255) not null,
    email varchar(255) unique not null,
    password varchar(255) not null,
    phone varchar(32),
    role user_role not null,
    foreign key (cart_id) references public.cart(id) on delete cascade
);

create type product_category as enum ('Electronic', 'Fashion', 'Home', 'Health', 'Sport', 'Books');
create table public.product (
     id uuid primary key,
     name varchar(255) not null,
     description text not null,
     price bigint not null,
     category product_category not null,
     photo_url text
);

create table public.cart_product (
    id uuid primary key,
    cart_id uuid not null,
    product_id uuid not null,
    quantity bigint not null,
    foreign key (cart_id) references public.cart(id) on delete cascade,
    foreign key (product_id) references public.product(id) on delete cascade,
    constraint uc_cart_product unique (cart_id,product_id)
);

create table public.shop (
     id uuid primary key,
     seller_id uuid not null,
     name varchar(255) not null,
     description text not null,
     requisites text not null,
     email varchar(255) unique not null,
     foreign key (seller_id) references public.user(id) on delete cascade
);

create table public.shop_product (
     id uuid primary key,
     shop_id uuid not null,
     product_id uuid not null,
     quantity bigint not null,
     foreign key (shop_id) references public.shop(id) on delete cascade,
     foreign key (product_id) references public.product(id) on delete cascade,
     check (quantity >= 0),
     constraint uc_shop_product unique (shop_id,product_id)
);

create type withdraw_status as enum ('Start', 'Ready', 'Done');
create table public.withdraw (
     id uuid primary key,
     shop_id uuid not null,
     comment text not null,
     sum bigint not null,
     status withdraw_status not null,
     foreign key (shop_id) references public.shop(id) on delete cascade
);

create table public.order_customer (
     id uuid primary key,
     customer_id uuid not null,
     address text not null,
     created_at timestamp not null,
     total_price bigint not null,
     payed boolean not null,
     foreign key (customer_id) references public.user(id) on delete cascade
);

create type order_shop_status as enum ('Start', 'Ready', 'Done');
create table public.order_shop (
     id uuid primary key,
     shop_id uuid not null,
     order_customer_id uuid not null,
     status order_shop_status not null,
     notified boolean not null,
     foreign key (shop_id) references public.shop(id) on delete cascade,
     foreign key (order_customer_id) references public.order_customer(id) on delete cascade,
     constraint uc_order_shop unique (shop_id,order_customer_id)
);

create table public.order_shop_product (
     id uuid primary key,
     order_shop_id uuid not null,
     product_id uuid not null,
     quantity bigint not null,
     foreign key (order_shop_id) references public.order_shop(id) on delete cascade,
     foreign key (product_id) references public.product(id) on delete cascade,
     constraint uc_order_shop_product unique (order_shop_id,product_id)
);


