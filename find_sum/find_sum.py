target = None
numbers = {}

with open('input.txt') as f:
    target = int(f.readline())
    for n in f.readline().split():
        num = int(n)
        numbers[num] = numbers.setdefault(num, 0) + 1

        
result = 0

for i in numbers:
    rest = target - i
    
    if i == rest and numbers[i] > 1:
        result = 1
    elif i != rest in numbers:
        result = 1


with open('output.txt', 'w') as f:
    f.write(str(result))

