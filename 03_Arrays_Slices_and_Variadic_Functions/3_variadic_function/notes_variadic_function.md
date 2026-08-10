What is a variadic function?
Functions in general accept only a fixed number of arguments. A variadic function is a function that accepts a variable number of arguments. If the last parameter of a function definition is prefixed by ellipsis …, then the function can accept any number of arguments for the last parameter.

Only the last parameter of a function can be variadic. We will learn why this is the case in the next section of this tutorial.

Syntax
func hello(a int, b ...int) {
}
In the above function, the parameter b is variadic since its type ...int is prefixed by ellipsis and it can accept any number of arguments.

=====================================================================================

By now I guess you would have understood why the variadic parameter should only be at the last.

Let’s try to make the first parameter of the hello function variadic.

The syntax will look like

func hello(b ...int, a int) {
}
In the above function, it is not possible to pass arguments to the parameter a because whatever argument we pass will be assigned to the first parameter b since it’s variadic. Hence variadic parameters can only be present at the last in the function definition. The above function will fail to compile with error can only use ... with final parameter in list

================================================================================

************************Append is a variadic function************************

Have you ever wondered how the append function in the standard library used to append values to a slice accepts any number of arguments. It’s because it’s a variadic function.

func append(slice []Type, elems ...Type) []Type
The above is the definition of append function. In this definition elems is a variadic parameter. Hence append can accept a variable number of arguments.

