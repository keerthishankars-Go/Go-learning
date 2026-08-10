A method is just a function with a special receiver type between the func keyword and the method name. The receiver can either be a struct type or non-struct type.

The syntax of a method declaration is provided below.

func (t Type) methodName(parameter list) {
}
The above snippet creates a method named methodName with receiver type Type. t is called as the receiver and it can be accessed within the method.

**************************Methods vs Functions**********************************

So why do we have methods when we can write the same program using functions. There are a couple of reasons for this. Let’s look at them one by one.

Go is not a pure object-oriented programming language and it does not support classes. Hence methods on types are a way to achieve behaviour similar to classes. Methods allow a logical grouping of behaviour related to a type similar to classes. In the above sample program, all behaviours related to the Employee type can be grouped by creating methods using Employee receiver type. For example, we can add methods like calculatePension, calculateLeaves and so on.

Methods with the same name can be defined on different types whereas functions with the same names are not allowed. Let’s assume that we have a Square and Circle structure. It’s possible to define a method named Area on both Square and Circle.
