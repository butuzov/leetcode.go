# #0848. [Shifting Letters](https://leetcode.com/problems/shifting-letters/?tab=Description) 

We have a string S of lowercase letters, and an integer array shifts.

Call the shift of a letter, the next letter in the alphabet, (wrapping around so that &amp;#39;z&amp;#39; becomes &amp;#39;a&amp;#39;).&amp;nbsp;

For example, shift(&amp;#39;a&amp;#39;) = &amp;#39;b&amp;#39;, shift(&amp;#39;t&amp;#39;) = &amp;#39;u&amp;#39;, and shift(&amp;#39;z&amp;#39;) = &amp;#39;a&amp;#39;.

Now for each shifts[i] = x, we want to shift the first i+1&amp;nbsp;letters of S, x times.

Return the final string&amp;nbsp;after all such shifts to S are applied.

Example 1:


Input: S = &amp;quot;abc&amp;quot;, shifts = [3,5,9]
Output: &amp;quot;rpl&amp;quot;
Explanation: 
We start with &amp;quot;abc&amp;quot;.
After shifting the first 1 letters of S by 3, we have &amp;quot;dbc&amp;quot;.
After shifting the first 2 letters of S by 5, we have &amp;quot;igc&amp;quot;.
After shifting the first 3 letters of S by 9, we have &amp;quot;rpl&amp;quot;, the answer.


Note:


	1 &amp;lt;= S.length = shifts.length &amp;lt;= 20000
	0 &amp;lt;= shifts[i] &amp;lt;= 10 ^ 9