tests = [
    # (input, target, output)
    ('', 1, 0),
    ('0', 0, 0),
    ('1 1 1 1', 2, 1),
    ('1 1 1 1', 3, 0),
    ('1 7 3 4 7 9', 13, 1),
    ('1 7 3 4 7 9', 3, 0),
    ('1', 0, 0),
    ('1', 1, 0),    
    ('999999999 ' * 1_000_000 + '7 4', 11, 1),
    ('999999999 ' * 1_000_000, 10, 0),
    ('999999999 ' * 1_000_000, 999999999 * 2, 1)]


from subprocess import call

for i, (inp, target, output) in enumerate(tests):
    print('Test', i)

    with open('input.txt', 'w') as f:
        f.write(str(target) + '\n')
        f.write(str(inp))
        
    call(['python3', 'find_sum.py'])

    result = None
    with open('output.txt') as f:
        result = int(f.readline())

    print ('Test', i, output == result)
    
