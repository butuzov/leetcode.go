# Power of Two

  Given an integer, write a function to determine if it is a power of two.

## Solution for O_(log n)_  
  ```pseudocode
  // starting with power = 1 (1 is power 0 for any signed int base )
  power = 1;
  while True {
    // n is smaller than power of 2 in current itteration.
    // returning false and existing procedure.
    if n < power {
      return false
    }
    // n is equivalent to our power variable
    // returning true and existing function  
    if n == power {
      return true;
    }
    power = power * 2;
  }
  ```

## Follow up

  Could you solve it without loops/recursion?

## Solution for O_(1)_  

## Given an integer, write a function to determine if it is a power of two.

  * https://code.tutsplus.com/articles/understanding-bitwise-operators--active-11301
  * http://php.net/manual/ru/language.operators.bitwise.php
  * https://wiki.python.org/moin/BitwiseOperators
  * https://en.wikipedia.org/wiki/Power_of_two
