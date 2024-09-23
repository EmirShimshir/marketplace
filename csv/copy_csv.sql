\copy public.cart (id, price) from './csv/cart.csv' delimiter ',' csv header;
\copy public.user (id, name, surname, email, password, phone, cart_id, role) from './csv/user.csv' delimiter ',' csv header NULL as 'NULL';
\copy public.product (id, name, description, price, category, photo_url) from './csv/product.csv' delimiter ',' csv header NULL as 'NULL';
\copy public.shop (id, seller_id, name, description, requisites, email) from './csv/shop.csv' delimiter ',' csv header NULL as 'NULL';
\copy public.shop_product (id, shop_id, product_id, quantity) from './csv/shop_product.csv' delimiter ',' csv header NULL as 'NULL';
