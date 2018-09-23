# #0441. [Arranging Coins](https://leetcode.com/problems/arranging-coins/#/description)

You have a total of _n_ coins that you want to form in a staircase shape, where every _k_-th row must have exactly _k_ coins.
 
Given _n_, find the total number of _full_ staircase rows that can be formed.

`n is a non-negative integer and fits within the range of a 32-bit signed integer.

Example 1: `n = 5`
------------------

The coins can form the following rows:
```
¤
¤ ¤
¤ ¤
```

Because the 3rd row is incomplete, we return `2`.

Example 2: `n = 8`
------------------


The coins can form the following rows:

```
¤
¤ ¤
¤ ¤ ¤
¤ ¤
```

Because the 4th row is incomplete, we return `3`.

---


### Solitions_1 - O ( log n )

simple count how many coins we can use to build stairs. 

-----------
  ```
  stair = 0
  stairs = 0
  while coins > 0 
    // incrementing stair 
    if coins < stair
      break loop
     
    // decrementing coins number
    coins = coins - stair
    // incrementing stair
    stairs = stairs + 1
      
  return coins.
 ```
 
### Solition_2 O(1)  

It really looks like its O(1) complexity, but it's not. Square root isn't a basic arithmetic operation computer can do. It's looks more like O (log n), due square root calculations... 

But still its faster then solution-1


  So we need to build a sequense of stairs that will fit into n coins.
  
  `1 + 2 + 3 + 4 + x = n`
  This called arithmetic sequense ([arithmetic progression](https://www.mathsisfun.com/algebra/sequences-sums-arithmetic.html) as I was told in school)
  
  Formula for sum of ap (arithmetic progression) is 
  
  ```
  2*1 + 1 * (K-1) 
  --------------- * K <= N
  2
  ```
  
  Simplified...
  
  `2  + (K - 1) * K <= 2N`
  
  Simplified...
  
  `2K + K^K - K     <= 2N`
  
  Simplified...
  
  `K^K + K          <= 2N`
  
  Simplified...
  
  `k^k + K - 2N     <= 0`
  
  So `quadratic equation` ...
  
  `a*x^x + b*x + c = 0`
  
  We need to find `discriminant` for it in order to solve equation and formula for `discriminant` is 
  
  `b^b – 4ac`
  
  and formula for x becomes 
  
  `(-b + square_root_of ( D ) ) / 2*a`
  
  so it becomes 
  
  ` -b + square_root_of( b**b - 4*a*c) ) / 2*a` 
  
  our data is a: 1, b:1, c:-2n
  `x = floar( -1 + square_root_of( 1 + 8*n ) / 2 ) `
