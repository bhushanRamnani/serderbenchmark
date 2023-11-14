import argparse
import json
import random
import string
from dataclasses import dataclass, asdict
import time
import sys

@dataclass
class Address:
    street: str
    city: str
    state: str
    zip_code: int

@dataclass
class Movie:
    name: str
    year: int
    director: str
    producer: str
    actors: list
    genre: str

@dataclass
class Person:
    first_name: str
    last_name: str
    address: Address
    phone_number: int
    favorite_movies: list[Movie]

def generate_random_string(length=10):
    letters = string.ascii_letters
    return ''.join(random.choice(letters) for _ in range(length))

def generate_random_number():
    return random.randint(10000, 99999)

def generate_random_movie():
    return Movie(
        name=generate_random_string(),
        year=random.randint(1980, 2022),
        director=generate_random_string(),
        producer=generate_random_string(),
        actors=[generate_random_string() for _ in range(3)],
        genre=generate_random_string()
    )

def generate_random_address():
    return Address(
        street=generate_random_string(),
        city=generate_random_string(),
        state=generate_random_string(),
        zip_code=generate_random_number()
    )

def generate_random_person():
    return Person(
        first_name=generate_random_string(),
        last_name=generate_random_string(),
        address=generate_random_address(),
        phone_number=generate_random_number(),
        favorite_movies=[generate_random_movie() for _ in range(20)]
    )

def measure_serialization_time(persons, iterations):
    total_serialization_time = 0

    for _ in range(iterations):
        start_time_serialization = time.time()
        persons_json_string = json.dumps([asdict(person) for person in persons], indent=2)
        serialization_time = (time.time() - start_time_serialization) * 1000
        total_serialization_time += serialization_time

    average_serialization_time = total_serialization_time / iterations
    return average_serialization_time

def measure_deserialization_time(persons_json_string, iterations):
    total_deserialization_time = 0

    for _ in range(iterations):
        start_time_deserialization = time.time()
        deserialized_persons = [Person(**item) for item in json.loads(persons_json_string)]
        deserialization_time = (time.time() - start_time_deserialization) * 1000
        total_deserialization_time += deserialization_time

    average_deserialization_time = total_deserialization_time / iterations
    return average_deserialization_time

def main():
    parser = argparse.ArgumentParser(description="Serialize and deserialize Person objects.")
    parser.add_argument('-i', type=int, default=10, help='Number of iterations for serialization and deserialization')
    parser.add_argument('-p', type=int, default=100, help='Number of Person objects to generate')

    args = parser.parse_args()
    print(f"Number of persons: {args.p}. Iterations: {args.i}")

    # Generate random Person objects
    persons = [generate_random_person() for _ in range(args.p)]

    # Measure serialization time
    average_serialization_time = measure_serialization_time(persons, args.i)
    print(f"Avg Serialization time ({args.i} iterations): {average_serialization_time:.6f} milliseconds")

    # Serialize to JSON-formatted string
    persons_json_string = json.dumps([asdict(person) for person in persons], indent=2)
    
    # Print the size of the serialized JSON string
    serialized_data_size = sys.getsizeof(persons_json_string)
    print(f"Size of serialized data: {serialized_data_size} bytes")

    # Measure deserialization time
    average_deserialization_time = measure_deserialization_time(persons_json_string, args.i)
    print(f"Avg Deserialization time ({args.i} iterations): {average_deserialization_time:.6f} milliseconds")


if __name__ == '__main__':
    main()
