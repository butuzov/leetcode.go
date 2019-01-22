# #0179. [Largest Number](https://leetcode.com/problems/largest-number/?tab=Description)

Given a list of non negative integers, arrange them such that they form the largest number.

For example, given [3, 30, 34, 5, 9], the largest formed number is 9534330.

Note: The result may be very large, so you need to return a string instead of an integer.

Credits: Special thanks to @ts for adding this problem and creating all test cases.


  A  |  B  | CmpA_1 | vs  | CmpB_1  | CmpA_2 | vs | CmpB_2 |Combined        
-----|-----|--------|-----|---------|--------|----|--------|----------------
  1  | 100 | 100    |  <  | 100     |        |    |        | 1_100 vs 100_1
  2  | 100 | 200    |  <  | 100     |        |    |        | 2_100 vs 100_2
 100 | 3   | 100    |  >  | 300     |        |    |        | 3_100 vs 100_3
 121 | 12  | 121    |  >  | 12[0]   |  121   | >  | [1]12  | 12_121 vs 121_12
