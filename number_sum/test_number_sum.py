from collections import namedtuple
from subprocess import PIPE, run

T = namedtuple('T', 'a b sum')

tests = [
    T(0, 0, 0),
    T(0, 1, 1),
    T(1, 0, 1),
    T(100, 11, 111),
    T(9999999999999999999999999, 999999999999999, 9999999999999999999999999 + 999999999999999),
    T(10**300, 10**500, 10**300 + 10**500),
    T(int('9' * 1000), int('8' * 1000), int('9' * 1000) + int('8' * 1000))]

cmd = ['./a.out']
    
for i, test in enumerate(tests):
    print('Test', i)
    result = run(cmd, input=f'{test.a} {test.b}',
                 encoding='ascii', stdout=PIPE)
    print('Test', i, test.sum == int(result.stdout))

print('Done')
