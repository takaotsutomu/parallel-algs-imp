# Sorting
## Parallel Merge Sort
Implemented in Go, so you can play with it on *The Go Playground*: [***Parallel Merge Sort***](https://go.dev/play/p/VwQMPd4KLNq).

### The algorithm for merge
**Input:** Two sorted arrays $A$, $B$ each of length $n$
**Output:** Merged array $M$, consisting of elements of $A$ and $B$ in sorted order

**for** each $x\in A$ **do in parallel** 

- Do a binary insertion search to find where $x$ would be added in $B$, 
- The final rank of $x$ given by $\text{rank}_M(x) = \text{rank}_A(x) + \text{rank}_B(x)$. 

**end** 
Do the symmetric for all $x\in B$. 

### The rationale behind the algorithm
For an element $x$ in $A$, we know how many elements (say $a$) in $A$ come before $x$ since we have sorted $A$. In order to know the rank of the element in the merged array $M$, we have to know how many elements (say $b$) in the other sorted array $B$ are less than $x$. Once we know $b$, then we know we should place $x$ in the $(a+b)^{th}$ position in the merged array $M$. We find $b$ by performing a binary insertion search over $B$. 
