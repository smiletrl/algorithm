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
H_{n} = (a_{1}b^{n-1} + a_{2}b^{n-2} + ... + a_{n-1}b + a_{n}) \% m
$$

I'm going to do some simple conversions to get the rolling hash function.

## Conversion 1

Given

$$
a = i_{a}m + j_{a}
$$

We get

$$
H = (a+b) \% m
$$

$$
H = (i_{a}m + j_{a} + b) \%m
$$

$$
H = (j_{a} + b) \%m
$$

$$
H = (a \% m + b) \%m
$$

## Conversion 2

Given

$$
a = i_{a}m + j_{a}
$$

We get

$$
H = (ab) \% m
$$

$$
H = (i_{a}m + j_{a})b \%m
$$

$$
H = (i_{a}mb + j_{a}b) \%m
$$

$$
H = (j_{a}b) \%m
$$

$$
H = ((a \%m) b) \%m
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
H_{n} = (x + y) \% m
$$

$$
H_{n} = (x \% m + y \% m) \% m
$$

$$
(H_{n} + m) \% m = (x \% m + y \% m) \% m
$$

$$
(H_{n} + m - x \% m) \% m = (y \% m) \% m
$$

$$
(H_{n} + m - x \% m) \% m = y \% m
$$

### Part two

$$
H_{n+1} = (yb + z) \% m
$$

$$
H_{n+1} = (yb \% m + z \% m) \% m
$$

$$
H_{n+1} = (((y \% m) b) \% m + z \% m) \% m
$$

$$
H_{n+1} = ((((H_{n} + m - x \% m) \% m) b) \% m + z \% m) \% m
$$

$$
H_{n+1} = (((H_{n} + m - x \% m) \% m) b + z) \% m
$$

This is what function `func (p *Path) rollingHash()` is doing.
