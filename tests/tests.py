""" 
1. Build the binary
2. Test all code samples in code folder
3. Any that has an exit code of not zero, the test didn't pass
"""

import subprocess
import os
import sys

TEST_FILE_NAMES = ["primes", "factorial", "fibonacci"]

BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
EXECUTABLE = os.path.join(BASE_DIR, "main.go")

BUILD_PATH = os.path.join(BASE_DIR, "bin")
OUTPUT_NAME = "lang"

BINARY_NAME = f"{BUILD_PATH}/{OUTPUT_NAME}"

CODE_PATH = os.path.join(BASE_DIR, "code")

output = subprocess.run(["go", "build", "-o", BINARY_NAME, EXECUTABLE])

if output.returncode != 0:
    print("Failed to create go binary")
    sys.exit(1)

for file_name in TEST_FILE_NAMES:
    print(
        f"\n====================== Executing {file_name} ================================"
    )

    execution = subprocess.run([BINARY_NAME, os.path.join(CODE_PATH, file_name)])

    print(
        f"==================== Finished Executing {file_name} =========================\n"
    )
