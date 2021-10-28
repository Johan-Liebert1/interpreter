import os

IGNORE_DIRS = ["testing"]

BASE_DIR = os.path.dirname(os.path.abspath(__file__))


def count_lines(directory, depth=0):
    total_lines = 0

    directory = os.path.join(BASE_DIR, directory)

    for dir in os.listdir(directory):
        if not os.path.isfile(dir) and dir in IGNORE_DIRS:
            continue

        complete_path = os.path.join(directory, dir)

        if os.path.isdir(dir):
            total_lines += count_lines(complete_path, depth + 1)

        elif dir.endswith(".go"):
            with open(complete_path) as file:
                total_lines += len(file.readlines())

    return total_lines
