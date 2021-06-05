# charcount

what? it counts char frequency recursively in all files `.go`.

why? it would be nice to know such thing when you want to decide which keyboard layout to pick (in pretty damn seriously way 🤣).

go? yeah, `golang` is my language now. but it's not that difficult to make a regex filter to match whatever file you want. you got the power.

---

### install

`go install github.com/namp10010/charcount`

### run

`charcount <dir>`

example running agains this project

`charcount .`

```bash
$ charcount .
main.go
space: 9.08%
r: 7.08%
tab: 6.83%
t: 5.58%
e: 5.50%
nl: 5.50%
...
```

running against a huge code base, funny enough even in the code `e` still the king

```text
space: 17.36%
e: 5.42%
0: 4.49%
t: 3.98%
r: 3.72%
n: 3.18%
tab: 3.02%
i: 2.94%
x: 2.93%
nl: 2.90%
a: 2.81%
o: 2.59%
,: 2.56%
s: 2.48%
l: 1.89%
c: 1.80%
u: 1.52%
f: 1.47%
d: 1.44%
p: 1.30%
:: 1.07%
.: 1.02%
1: 0.99%
_: 0.96%
E: 0.96%
m: 0.95%
/: 0.95%
): 0.93%
(: 0.93%
S: 0.85%
T: 0.85%
=: 0.84%
g: 0.79%
b: 0.76%
h: 0.76%
2: 0.70%
3: 0.70%
4: 0.65%
I: 0.64%
y: 0.62%
6: 0.61%
8: 0.59%
C: 0.59%
": 0.58%
{: 0.57%
}: 0.57%
A: 0.56%
R: 0.53%
O: 0.51%
v: 0.49%
N: 0.47%
P: 0.47%
D: 0.47%
9: 0.44%
5: 0.42%
7: 0.41%
F: 0.37%
L: 0.36%
M: 0.33%
k: 0.31%
w: 0.29%
B: 0.26%
\: 0.23%
*: 0.21%
H: 0.21%
U: 0.20%
G: 0.20%
[: 0.16%
]: 0.16%
Y: 0.16%
V: 0.16%
-: 0.15%
K: 0.10%
W: 0.10%
!: 0.10%
&: 0.08%
;: 0.08%
z: 0.08%
q: 0.08%
X: 0.08%
+: 0.08%
j: 0.06%
`: 0.05%
': 0.05%
<: 0.05%
|: 0.04%
>: 0.03%
%: 0.03%
Q: 0.02%
J: 0.02%
Z: 0.01%
â: 0.01%
Ù: 0.01%
```

### todo

* maybe make it case insensitive?