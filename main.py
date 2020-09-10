import csv
import re

HAS_NUMERICS = re.compile(r"^.*\d+.*$")
INT_REGEX = re.compile(r"^\s*\d+\s*$")
FLOAT_REGEX = re.compile(r"^\s*\d*\.?\d*\s*$")


def try_parse_number(number_thing):
    if type(number_thing) != str:
        return number_thing

    if not HAS_NUMERICS.match(number_thing):
        return number_thing

    if INT_REGEX.match(number_thing):
        return int(number_thing)

    if FLOAT_REGEX.match(number_thing):
        return float(number_thing)

    return number_thing


def read_csv(file_path):
    with open(file_path, newline='') as csvfile:
        reader = csv.reader(csvfile, delimiter=';')
        data = []
        for row in reader:
            data.append([try_parse_number(x) for x in row])
            # yield map(try_parse_number, row)
        return data


if __name__ == '__main__':
    data = read_csv('dataset.csv')

    print(len(data))
    with open("out.csv", "w", newline="") as f:
        writer = csv.writer(f)
        writer.writerows(data)
