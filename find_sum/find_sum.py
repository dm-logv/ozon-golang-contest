target = None
numbers = {}
result = 0

with open('input.txt') as f:
    target = int(f.readline())
    for n in f.readline().split():
        num = int(n)
        numbers[num] = numbers.setdefault(num, 0) + 1

        rest = target - num

        if (num == rest and numbers[num] > 1
            or num != rest and rest in numbers):
            result = 1
            break


with open('output.txt', 'w') as f:
    f.write(str(result))

