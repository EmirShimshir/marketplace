create table public.cart (
    id uuid,
    price bigint
);

create type user_role as enum ('Customer', 'Seller', 'Moderator');
create table public.user (
    id uuid,
    cart_id uuid,
    name varchar(255),
    surname varchar(255),
    email varchar(255),
    password varchar(255),
    phone varchar(32),
    role user_role
);

create type product_category as enum ('Electronic', 'Fashion', 'Home', 'Health', 'Sport', 'Books');
create table public.product (
    id uuid,
    name varchar(255),
    description text,
    price bigint,
    category product_category,
    photo_url text
);

create table public.cart_product (
    id uuid,
    cart_id uuid not null,
    product_id uuid not null,
    quantity bigint not null
);

create table public.shop (
    id uuid,
    seller_id uuid,
    name varchar(255),
    description text,
    requisites text,
    email varchar(255)
);

create table public.shop_product (
    id uuid,
    shop_id uuid,
    product_id uuid,
    quantity bigint
);

create type withdraw_status as enum ('Start', 'Ready', 'Done');
create table public.withdraw (
    id uuid,
    shop_id uuid,
    comment text,
    sum bigint,
    status withdraw_status
);

create table public.order_customer (
    id uuid,
    customer_id uuid,
    address text,
    created_at timestamp,
    total_price bigint,
    payed boolean
);

create type order_shop_status as enum ('Start', 'Ready', 'Done');
create table public.order_shop (
    id uuid,
    shop_id uuid,
    order_customer_id uuid,
    status order_shop_status,
    notified boolean
);

create table public.order_shop_product (
    id uuid,
    order_shop_id uuid,
    product_id uuid,
    quantity bigint
);
