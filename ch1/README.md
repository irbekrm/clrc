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
