from sys import stderr, stdin


read_numbers = set()

for row in stdin:
    number = int(row)

    if number in read_numbers:
        read_numbers.remove(number)
        continue

    read_numbers.add(number)

if len(read_numbers) != 1:
    print('Something wrong with input data. Numbers without pairs:',
          ', '.join(map(str, read_numbers)), file=stderr)
else:
    print(read_numbers.pop())

