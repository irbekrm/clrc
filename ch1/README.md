Exercise 2.1-3

for j = 1 to A.length
  if A[i] == v
    return i
return nil

Initialization: i = 1 (1st index of A)
Maintenance: value `v` is not within the values A[0..i-1]
Termination: A[i] == v or i = A.length + 1 