from math import sqrt
from itertools import count, islice

def isPrime(n):
	return n > 1 and all(n % i for i in islice(count(2), int(sqrt(n) - 1)))

# Sorry, no generic solution today
sum = 0

for i in range(109900, 126901, 17):
	if not isPrime(i):
		sum += 1

print("sum:", sum)
