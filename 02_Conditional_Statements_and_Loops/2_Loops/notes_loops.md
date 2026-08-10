A loop is used to execute a block of code repeatedly until a condition is satisfied.

for is the only loop available in Go. Go doesn’t have while or do while loops which are present in other languages like C.

===============================================

for loop syntax

for initialisation; condition; post {
}

The initialisation statement will be executed only once. After the loop is initialised, the condition is checked. 
If the condition evaluates to true, the body of the loop inside the { } will be executed followed by the post statement. 
The post statement will be executed after each successful iteration of the loop. After the post statement is executed, the condition will be rechecked. If it’s true, the loop will continue executing, else the for loop terminates.

All the three components namely initialisation, condition and post are optional in Go. Let’s look at an example to understand for loop better.







