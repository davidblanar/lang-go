### An interpreter for a made up programming language written in GO

This is a re-implementation of a similar interpreter written in JS.\
For the original project, see [here](https://github.com/davidblanar/lang).

To build an executable:\
``go build``

To run the interpreter on a file:\
``./lang-go -f <path_to_file>``


Features:
<br/>
- [x] Numbers
- [x] Strings
- [x] Booleans
- [x] Null value
- [x] Comparison operators (=, >, <, >=, <=)
- [x] Basic arithmetic operations (+, -, *, /, %)
- [x] Variable declaration and reference
- [x] Function declaration and calls with function-scoped variables
- [x] Conditionals
- [x] Comments
- [x] Arrays
- [x] Objects (dictionaries)

For examples see ``src/files/``
