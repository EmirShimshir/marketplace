from faker import Faker
import uuid
import csv
import faker_commerce

fake = Faker()
fake.add_provider(faker_commerce.Provider)


def create_csv_title(filename, data):
    with open(filename, 'w', newline='', encoding='utf-8') as file:
        writer = csv.writer(file)
        writer.writerow(data)


def write_to_csv(filename, data):
    with open(filename, 'a', newline='', encoding='utf-8') as file:
        writer = csv.writer(file)
        writer.writerow(data)


def create_products(shop_id, count, product_path, shop_product_path):
    create_csv_title(product_path, [
        "id",
        "name",
        "description",
        "price",
        "category",
        "photo_url",
    ])
    create_csv_title(shop_product_path, [
        "id",
        "shop_id",
        "product_id",
        "quantity",
    ])
    for num in range(1, count + 1):
        product_id = str(uuid.uuid4())
        name = fake.ecommerce_name() + str(num)
        description = fake.ecommerce_material()
        price = fake.ecommerce_price()
        category = fake.random_element(['Electronic', 'Fashion', 'Home', 'Health', 'Sport', 'Books'])
        photo_url = "http://photo.url"
        write_to_csv(product_path, [
            product_id,
            name,
            description,
            price,
            category,
            photo_url,
        ])
        shop_product_id = str(uuid.uuid4())
        quantity = fake.random_int(1, 10000)
        write_to_csv(shop_product_path, [
            shop_product_id,
            shop_id,
            product_id,
            quantity,
        ])


def main():
    shop_id = "30e18bc1-4354-4937-9a3b-03cf0b7027b1"
    count = 10000
    product_path = "csv/product.csv"
    shop_product_path = "csv/shop_product.csv"

    create_products(shop_id, count, product_path, shop_product_path)


if __name__ == "__main__":
    main()
