target = None
numbers = None

with open('input.txt') as f:
    target = int(f.readline())
    numbers = [int(n) for n in f.readline().split()]

result = 0

for i, n in enumerate(numbers):
    for j in range(i + 1, len(numbers)):
        if n + numbers[j] == target:
            result = 1
            break
    else:
        continue
    break

with open('output.txt', 'w') as f:
    f.write(str(result))

