# #0012. [Integer to Roman](https://leetcode.com/problems/integer-to-roman/?tab=Description)

Roman numerals are represented by seven different symbols: `I, V, X, L, C, D and M`.

```
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```

For example : two is written as `II` in Roman numeral, just two one's added together. Twelve is written as, `XII`, which is simply `X + II`. The number twenty seven is written as `XXVII`, which is `XX + V + II`.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:


	I can be placed before V (5) and X (10) to make 4 and 9.&amp;nbsp;
	X can be placed before L (50) and C (100) to make 40 and 90.&amp;nbsp;
	C can be placed before D (500) and M (1000) to make 400 and 900.


Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

Example 1:


Input:&amp;nbsp;3
Output: &amp;quot;III&amp;quot;

Example 2:


Input:&amp;nbsp;4
Output: &amp;quot;IV&amp;quot;

Example 3:


Input:&amp;nbsp;9
Output: &amp;quot;IX&amp;quot;

Example 4:


Input:&amp;nbsp;58
Output: &amp;quot;LVIII&amp;quot;
Explanation: C = 100, L = 50, XXX = 30 and III = 3.


Example 5:


Input:&amp;nbsp;1994
Output: &amp;quot;MCMXCIV&amp;quot;
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
