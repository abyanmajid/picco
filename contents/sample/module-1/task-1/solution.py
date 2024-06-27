import sys

def add(a, b):
    return a + b

def subtract(a, b):
    return a - b

def multiply(a, b):
    return a * b

def divide(a, b):
    if b == 0:
        return "Error: Division by zero"
    return a / b

calculation_type = sys.argv[1]
num1 = int(sys.argv[2])
num2 = int(sys.argv[3])

if calculation_type == 'add':
    print(add(num1, num2))
elif calculation_type == 'subtract':
    print(subtract(num1, num2))
elif calculation_type == 'multiply':
    print(multiply(num1, num2))
elif calculation_type == 'divide':
    print(divide(num1, num2))