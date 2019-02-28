Exercise 2.1-3

**Input**: A sequence of `n` numbers `A={a<sub>1</sub>, a<sub>2</sub>,..., a<sub>n</sub>}` and a value `v`

**Output**: An index `i` such that `v == A[i]` or `NIL`

LINEAR SEARCH
```
for j = 1 to A.length
  if A[i] == v
    return i
return NIL
```

**LOOP INVARIANT**:

**Initialization**: `i = 1`(_i_ holds the first index of _A_)

**Maintenance**: value `v` is not within the values `A[0..i-1]`

**Termination**: `A[i] == v` or `i == A.length + 1`


Exercise 2.1-4

**Input**: A sequence of binary digits `A={a<sub>1</sub>, a<sub>2</sub>,...,a<sub>n</sub>}`, a sequence of binary digits
`B={b<sub>1</sub>,b<sub>2</sub>,...b<sub>n</sub>}`

**Output**: A sequence of binary digits `C={c<sub>1</sub>,c<sub>2</sub>,...c<sub>n+1</sub>}`

BINARY ADDITION:

```
supplement = 0
let C[1..n+1] be a new array

for j = A.length downto 1
  sum = A[j] + B[j] + supplement
  if sum < 2
    C[j + 1] = sum
    supplement = 0
  else if sum == 2
    C[j + 1] = 0
    supplement = 1
  else // sum == 3
    C[j + 1] = 1
    supplement = 1
    
C[1] = supplement
```
