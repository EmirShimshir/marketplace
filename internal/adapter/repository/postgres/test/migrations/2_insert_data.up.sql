-- insert carts
insert into cart (id, price)
values ('30e18bc1-4354-4937-9a3b-03cf0b7034cc', 0);
insert into cart (id, price)
values ('30e18bc1-4354-4937-9a3b-03cf0b7034cd', 0);


-- insert users
insert into public.user (id, name, surname, email, password, phone, cart_id, role)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', 'Timur', 'Musin', 'hanoys@mail.ru', 'qwerty',  '+79992233555', '30e18bc1-4354-4937-9a3b-03cf0b7034cc', 'Customer');
insert into public.user (id, name, surname, email, password, cart_id, role)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cc', 'Emir', 'Shimshir', 'emir@gmail.com', '12345', '30e18bc1-4354-4937-9a3b-03cf0b7034cd', 'Customer');

-- insert products
insert into public.product (id, name, description, price, category, photo_url)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027a1', 'iphone 15', 'apple IOS', 129990, 'Electronic',  'photo/1.png');
insert into public.product (id, name, description, price, category, photo_url)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027a2', 'harry potter', 'Rouling', 2990, 'Books', 'photo/2.png');

-- insert cart_product
insert into public.cart_product (id, cart_id, product_id, quantity)
values ('30e18bc1-4354-4937-9a3b-03cf0b702aa1', '30e18bc1-4354-4937-9a3b-03cf0b7034cc', '30e18bc1-4354-4937-9a3b-03cf0b7027a1', 2);
insert into public.cart_product (id, cart_id, product_id, quantity)
values ('30e18bc1-4354-4937-9a3b-03cf0b702aa2', '30e18bc1-4354-4937-9a3b-03cf0b7034cc', '30e18bc1-4354-4937-9a3b-03cf0b7027a2', 1);

-- insert shop
insert into public.shop (id, seller_id, name, description, requisites, email)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027b1', '30e18bc1-4354-4937-9a3b-03cf0b7027cb', 'Apple Store', 'found 1998', 'Alabama', 'Apple@mail.ru');

-- insert shop_product
insert into public.shop_product (id, shop_id, product_id, quantity)
values ('30e18bc1-4354-4937-9a3b-03cf0b702ac1', '30e18bc1-4354-4937-9a3b-03cf0b7027b1', '30e18bc1-4354-4937-9a3b-03cf0b7027a1', 5);

-- insert withdraw
insert into public.withdraw (id, shop_id, comment, sum, status)
values ('30e18bc1-4354-4937-9a3b-03cf0b702ad1', '30e18bc1-4354-4937-9a3b-03cf0b7027b1', 'comment', 9999, 'Done');

-- insert order_customer
insert into public.order_customer (id, customer_id, address, created_at, total_price, payed)
values ('30e18bc1-4354-4937-9a3b-03cf0b702ae1', '30e18bc1-4354-4937-9a3b-03cf0b7027cc', 'Pushkina 1-2-3', '2022-10-10 11:30:30', 0, 'false');

-- insert order_shop
insert into public.order_shop (id, shop_id, order_customer_id, status, notified)
values ('30e18bc1-4354-4937-9a3b-03cf0b702ee1', '30e18bc1-4354-4937-9a3b-03cf0b7027b1', '30e18bc1-4354-4937-9a3b-03cf0b702ae1', 'Start', 'false');

-- insert order_shop_product
insert into public.order_shop_product (id, order_shop_id, product_id, quantity)
values ('30e18bc1-4354-4937-9a3b-03cf0b70eee1', '30e18bc1-4354-4937-9a3b-03cf0b702ee1', '30e18bc1-4354-4937-9a3b-03cf0b7027a1', 1);
