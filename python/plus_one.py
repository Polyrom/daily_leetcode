def plus_one(digits: list[int]) -> list[int]:
    complement = 1
    i = len(digits) - 1
    while complement != 0:
        if i == -1:
            return [1] + digits
        digits[i] += complement
        if digits[i] > 9:
            digits[i] = 0
            i -= 1
        else:
            complement = 0
    return digits

def plus_one_optimized(digits: list[int]) -> list[int]:
    for i in range(len(digits) - 1, -1, -1):
        if digits[i] == 9:
            digits[i] = 0
        else:
            digits[i] += 1
            return digits
    return [1] + digits

def plus_one_oneline(digits: list[int]) -> list[int]:
    return [int(d) for d in str(int(''.join(map(str, digits))) + 1)]

print('My solution')
num = [1, 2, 8]
print(f'num={num}; expected=[1, 2, 9], actual={plus_one(num)}')
num = [1, 2, 9]
print(f'num={num}; expected=[1, 3, 0], actual={plus_one(num)}')
num = [1, 9, 9]
print(f'num={num}; expected=[2, 0, 0], actual={plus_one(num)}')
num = [9, 9]
print(f'num={num}; expected=[1, 0, 0], actual={plus_one(num)}')

print('Leetcode solutions, brushed up, but same approach')
num = [1, 2, 8]
print(f'num={num}; expected=[1, 2, 9], actual={plus_one_optimized(num)}')
num = [1, 2, 9]
print(f'num={num}; expected=[1, 3, 0], actual={plus_one_optimized(num)}')
num = [1, 9, 9]
print(f'num={num}; expected=[2, 0, 0], actual={plus_one_optimized(num)}')
num = [9, 9]
print(f'num={num}; expected=[1, 0, 0], actual={plus_one_optimized(num)}')

print('Leetcode solutions, one-liner, conversion to string')
num = [1, 2, 8]
print(f'num={num}; expected=[1, 2, 9], actual={plus_one_oneline(num)}')
num = [1, 2, 9]
print(f'num={num}; expected=[1, 3, 0], actual={plus_one_oneline(num)}')
num = [1, 9, 9]
print(f'num={num}; expected=[2, 0, 0], actual={plus_one_oneline(num)}')
num = [9, 9]
print(f'num={num}; expected=[1, 0, 0], actual={plus_one_oneline(num)}')