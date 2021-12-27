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

running against a huge code base, funny enough even in the code `e` is still the king (ok, not including space)

|                                                       15 |                                                       30 |                                                       45 |                            60                                 |                                                              |
| -----------------------------------------------------------: | -----------------------------------------------------------: | -----------------------------------------------------------: | -----------------------------------------------------------: | -----------------------------------------------------------: |
| space: 12.13%<br/>e: 7.70%<br/>t: 5.54%<br/>r: 4.68%<br/>s: 4.11%<br/>o: 4.02%<br/>n: 3.99%<br/>i: 3.96%<br/>a: 3.83%<br/>0: 3.24%<br/>x: 2.94%<br/>,: 2.92%<br/>c: 2.62%<br/>nl: 2.58%<br/>l: 2.53% | tab: 2.43%<br/>p: 2.15%<br/>d: 2.13%<br/>u: 2.04%<br/>f: 1.79%<br/>m: 1.52%<br/>g: 1.45%<br/>/: 1.42%<br/>.: 1.36%<br/>h: 1.13%<br/>": 0.93%<br/>6: 0.89%<br/>b: 0.86%<br/>): 0.79%<br/>(: 0.79% | y: 0.76%<br/>_: 0.76%<br/>v: 0.71%<br/>1: 0.69%<br/>:: 0.68%<br/>2: 0.64%<br/>7: 0.53%<br/>=: 0.50%<br/>4: 0.50%<br/>{: 0.48%<br/>}: 0.48%<br/>3: 0.48%<br/>w: 0.45%<br/>5: 0.45%<br/>k: 0.35% | 9: 0.30%<br/>8: 0.29%<br/>\: 0.27%<br/>*: 0.27%<br/>q: 0.26%<br/>-: 0.19%<br/>j: 0.19%<br/>`: 0.17%<br/>[: 0.16%<br/>]: 0.16%<br/>z: 0.11%<br/>!: 0.08%<br/>&: 0.08%<br/>+: 0.06%<br/>;: 0.04% | ': 0.04%<br /> %: 0.03%<br /> <: 0.03% <br />>: 0.02%<br /> à: 0.02%<br /> \|: 0.02%<br /> ð: 0.02% ?: 0.01% |

for non-letter characters

|special|freq  |
|---|------|
|,| 2.92%|
|/| 1.42%|
|.| 1.36%|
|"| 0.93%|
|)| 0.79%|
|(| 0.79%|
|_| 0.76%|
|:| 0.68%|
|=| 0.50%|
|{| 0.48%|
|}| 0.48%|
|*| 0.27%|
|-| 0.19%|
|`| 0.17%|
|[| 0.16%|
|]| 0.16%|
|!| 0.08%|
|&| 0.08%|
|+| 0.06%|
|;| 0.04%

digits

|digit |freq  |
|---|------|
|0| 3.24%|
|6| 0.89%|
|1| 0.69%|
|2| 0.64%|
|7| 0.53%|
|4| 0.50%|
|3| 0.48%|
|5| 0.45%|
|9| 0.30%|
|8| 0.29%|
