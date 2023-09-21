# Real numbers (homework 1)

1. Implement a program that converts fractions between base 10 and base 2.

    1. Implement a function `func base2ToBase10(num string) (float64, error)`, which accepts
    a string representing a binary finite floating point number (e.g. `"11.001"`) and
    returns the corresponding floating point number in base 10. The function
    returns an error if the input is not correct.
    
    Implement tests for the function.

    2. 1. Implement a function `func base10ToBase2(num float64) (string, error)`, which accepts
    a floating point number and returns the string representing the corresponding
    binary floating point number. The returned number may contain a period.

    Implement tests for the function.

3. Write a program that solve the first task ("Pattern Matching") from
[this](https://codeforces.com/gym/104610/attachments) contest. 