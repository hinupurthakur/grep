# grep

`grep` functionality to be implemented in Golang using Cobra

### Basic Implementation: 
`grep <regex_to_searched> <file_name>`

O/P:
Print Matched Line (optional: highlight the matched word)
If no file exists then print: `grep: README.m: No such file or directory`


Test Driven Approach:
* Write test before writing the code.
* API Contract.


regex : args[0]
file_name: args[1]

utility function to read the file
then,
Match the regex and return the array of matched lines.

Test:
Failure: 
    if the file_name return `grep: README.m: No such file or directory`
    if the file exists then return lines based dummy regex



2 Ways:
1. Entire CLI Testing
2. Functionality

Funtional Core and Imperative Shell: If test first, then we test about the core of the application.

Directory Implementation:

grep foobar -r <dir> 

Print File path and the occurence lines
