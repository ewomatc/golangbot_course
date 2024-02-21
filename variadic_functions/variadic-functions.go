package main

func main() {
	hello(2, 3)

}

// Functions in general accept only a fixed number of arguments. A variadic function is a function that accepts a variable number of arguments. If the last parameter of a function definition is prefixed by ellipsis …, then the function can accept any number of arguments for that parameter.

// Only the last parameter of a function can be variadic. We will learn why this is the case in the next section of this tutorial.

// eXAMPLE
func hello(a int, b ...int) int { // b is variadic isnce it is the last parameyter followed by elipses '...'
	return a
}

// we can call the hello funnction like this
//hello(1, 2)
// and also like this
//hello(1, 2,, 4, 5, 7) // only the first argument is for, b has no limit to its number of args
// we can also call variadic functions without sny args,
// hello (1)
// we only pass an arg for a
