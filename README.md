## DESCRIPTION
Just a repo me practising DSA problems
I used neetcode [roadmap](https://neetcode.io/roadmap)

## PROBLEMS (Arrays and Hashing)

1. duplicate in array (set) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem1.go)
2. is Anagram (hmap, store character count) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem2.go)
3. two sums (hmap, key = ele, value = index, x + y = target i,e x = target - y) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem3.go)
4. k most frequent element (hmap , key = num, value = frequency, sorted bucket algo) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem4.go)
5. String Encode and Decode (attached length of string before actual string) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem5.go)
6. Products of Array Discluding Self (maintained product of left and right of an element) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem6.go)
7. longest consecutive sequence (set , if num - 1 is not in the set then num is the new sequence) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem7.go)

## PROBLEMS (Two Pointers)

8. is palindrome (compare two pointer one from start and one from end, skip if the byte/character is not alphanumeric) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem8.go)
9. Two Sum II - Input Array Is Sorted (array is sorted is ASC, left and right pointers) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem9.go)
10. Three Integer Sum (sort the array, positive sum cant be zero, impelement two sum(two pointer) logic in the inner loop) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem10.go)
11. Maximum Water Container ( used two pointer left and right , index will be the width and element is the actual height of bar, and find the area = min(height[left], height[right]) * (right - left) ) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem11.go)

## PROBLEMS (Stack)
12. Validate Parentheses [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem12.go)
13. Minimum Stack (maintained a minimum and normal stack, used minimum stack to  retrive the minimum value of the stack) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem13.go)
14. Evaluate Reverse Polish Notation [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem14.go)
15. Daily Temperatures (used stack to store the index of an temperatures, just start from reverse and think) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem15.go)

## PROBLEMS (Binary Search)
16. Binary Search [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem16.go)
17. Search 2D Matrix (find potential row by doing BS on first column and again do BS on that potential row) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem17.go)
18. Koko Eating Bananas (min and max speed of koko is 1 and max(piles), do that BS on that range and check againt the totaltime at that mid rate) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem18.go)
19. Find Minimum in Rotated Sorted Array (min elemnet will be always less than left and right) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem19.go)
20. Find Target in Rotated Sorted Array (first check target == mid, if not find the sorted and unsorted subarray and then check wether that target is fit on the sorted or unsorted subarray) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem20.go)
21. Time Based Key-Value Store (since set timestamp is in increasing order, do BS) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem21.go)

## PROBLEMS (Window Sliding)
22.  Maximum Average Subarray I (subtract the previous ele in left pointer and add the new element to the previous sum) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem22.go)
23.  Buy and Sell Crypto (dynamic size window) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem23.go)
24.  Longest Substring Without Duplicates (dynamic size window and set for checking repeated character) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem24.go)
25.  Minimum Window Substring (use left ptr to find the minimum sub string and right ptr to find that WS that contains the target sub string characters) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem25.go)
26.  Sliding Window Maximum (monotonic queue) [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem26.go)

## PROBLEMS (Link List)
27.  Reverse a singly linklist [mySolution](https://github.com/ShuaibKhan786/dsa/blob/main/problems/problem27.go)