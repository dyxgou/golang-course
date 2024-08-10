- Literals
Like some languages, golang assigns "zero values" to the built-in types, for example:

* Boolean : false
* String : ""
* int : 0


- Literals

A literal in go refersn to a writing out a number ort a string, there are four types of literals (and a werid one that is the fifth).

-- Literal Integers:

* base 10: The common integer literal base 10 (doesn't need a prefix)
* base 2: A literal bynary number is represented as `0b` -> 0b11
* base 8: The prefix for an octal number is `0o`, you can also use `0` but it tends to be confusing
* base 16: The prefix for an hexadecimal number is `0x`.

To make numbers easier to read you can put under scores between the digits.

`Example`: 5_614


-- Floating point Literals:

They can be decimal literals and can also have an positive or negative exponent.

`Example`: 6.03e23 -> 6.03 * 10 ^ 23

They can also be hexadecimals and have underscore between digits.


-- Rune literals:

These literals represent characters and are surrounded by single quoutes *THIS IS VERY IMPORTANT*

`Example`: 'a'

-- String Literal

There are many string backslash escaped runes to personalize your string literals, for example, there is:

* New line: \n:
* Tab: \t
* Quoute: \'
* Double Quoute: \"
* Backslash: \\

You can also make string literals that use all of the characters except the backquoute:

```go
const string : string = `This is a string
    literal`
```


-- Floats

The go floats use IEEE 754.

