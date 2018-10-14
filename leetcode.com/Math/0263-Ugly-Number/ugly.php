<?php 
// PHP program to find nth ugly number 
  
// This function divides a by  
// greatest divisible power of b 
function maxDivide($a, $b) 
{ 
    while ($a % $b == 0) 
    $a = $a / $b;  
    return $a; 
}  
  
// Function to check if a  
// number is ugly or not  
function isUgly($no) 
{ 
    $no = maxDivide($no, 2); 
    $no = maxDivide($no, 3); 
    $no = maxDivide($no, 5); 
      
    return ($no == 1)? 1 : 0; 
}  
  

var_dump(isUgly(14));

// Function to get the nth 
// ugly number 
// function getNthUglyNo($n) 
// { 
//     $i = 1;  
      
//     // ugly number count 
//     $count = 1;  
  
//     // Check for all integers  
//     // untill ugly count becomes n 
//     while ($n > $count) 
//     { 
//         $i++;      
//         if (isUgly($i)) 
//         $count++;  
//     } 
//     return $i; 
//     } 
  
//     // Driver Code 
//     $no = getNthUglyNo(150); 
//     echo "150th ugly no. is ". $no; 
  
// This code is contributed by Sam007 
?> 