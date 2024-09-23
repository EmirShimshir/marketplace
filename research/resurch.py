import csv

from locust import task, constant_throughput, HttpUser
import random


def get_column_as_array(filename, column_index):
    column_data = []
    with open(filename, 'r', newline='') as csvfile:
        reader = csv.reader(csvfile)
        for row in reader:
            if row[column_index] == "name":
                continue
            column_data.append(row[column_index])
    return column_data

column_array = get_column_as_array("./csv/product.csv", 1)


class CartUser(HttpUser):
    wait_time = constant_throughput(0.1)
    host = "http://localhost:8080"

    @task
    def add_to_cart(self):
        name = random.choice(column_array)
        self.client.get(f"/api/v1/product?name={name}")
