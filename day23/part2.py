b = 109900  # start
c = 126900  # end
h = 0  # counter

while True:
	f = 1
	d = 2
	while True:
		e = 2
		while True:
			if e * d - b == 0:
				f = 0
			e += 1
			if e - b == 0:
				break
		d += 1
		if d - b == 0:
			break
	if f == 0:
		h += 1  # if number is *not* prime
	if b - c == 0:
		break
	b += 17  # step

print(h)
