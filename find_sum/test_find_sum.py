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


from os import path, remove
from subprocess import call

inp_f = 'input.txt'
out_f = 'output.txt'

call('c++ -o find_sum find_sum.cpp'.split())

for i, (inp, target, output) in enumerate(tests):
    if path.exists(inp_f): remove(inp_f)
    if path.exists(out_f): remove(out_f)
    
    print('Test', i)

    with open(inp_f, 'w') as f:
        f.write(str(target) + '\n')
        f.write(str(inp))
        
    call(['./find_sum'])

    result = None
    with open(out_f) as f:
        result = int(f.readline())

    print ('Test', i, output == result)
    
