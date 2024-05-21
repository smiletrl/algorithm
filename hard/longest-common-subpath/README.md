# Solution

[Rabin–Karp_algorithm](https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm) is used here.

Let's assume the fixed length is n. To calulate hash value from 1 - n, the raw polynomial is

$$
H_{n} = a_{1}b^{n-1} + a_{2}b^{n-2} + ... + a_{n-1}b + a_{n}
$$

To calulate hash value from 2 - n+1, the raw polynomial is

$$
H_{n+1} = a_{2}b^{n-1} + a_{2}b^{n-2} + ... + a_{n}b + a_{n+1}
$$

Ideally,

$$
H_{n+1} = (H_{n} - a_{1}b^{n-1})b + a_{n+1}
$$

We want to add mod calculation to decrease number size

$$
H_{n} = (a_{1}b^{n-1} + a_{2}b^{n-2} + ... + a_{n-1}b + a_{n}) \bmod m
$$

I'm going to do some simple conversions to get the rolling hash function.

## Conversion 1

Given

$$
a = i_{a}m + j_{a}
$$

We get

$$
H = (a+b) \bmod m
$$

$$
H = (i_{a}m + j_{a} + b) \bmod m
$$

$$
H = (j_{a} + b) \bmod m
$$

$$
H = (a \bmod m + b) \bmod m
$$

## Conversion 2

Given

$$
a = i_{a}m + j_{a}
$$

We get

$$
H = (ab) \bmod m
$$

$$
H = (i_{a}m + j_{a})b \bmod m
$$

$$
H = (i_{a}mb + j_{a}b) \bmod m
$$

$$
H = (j_{a}b) \bmod m
$$

$$
H = ((a \bmod m) b) \bmod m
$$

## Conversion 3

### Part one

$$
x = a_{1}b^{n-1}
$$

$$
y = a_{2}b^{n-2} + ... + a_{n-1}b + a_{n}
$$

$$
z = a_{n+1}
$$

$$
H_{n} = (x + y) \bmod m
$$

$$
H_{n} = (x \bmod m + y \bmod m) \bmod m
$$

$$
(H_{n} + m) \bmod m = (x \bmod m + y \bmod m) \bmod m
$$

$$
(H_{n} + m - x \bmod m) \bmod m = (y \bmod m) \bmod m
$$

$$
(H_{n} + m - x \bmod m) \bmod m = y \bmod m
$$

### Part two

$$
H_{n+1} = (yb + z) \bmod m
$$

$$
H_{n+1} = (yb \bmod m + z \bmod m) \bmod m
$$

$$
H_{n+1} = (((y \bmod m) b) \bmod m + z \bmod m) \bmod m
$$

$$
H_{n+1} = ((((H_{n} + m - x \bmod m) \bmod m) b) \bmod m + z \bmod m) \bmod m
$$

$$
H_{n+1} = (((H_{n} + m - x \bmod m) \bmod m) b + z) \bmod m
$$

This is what function `func (p *Path) rollingHash()` is doing.
