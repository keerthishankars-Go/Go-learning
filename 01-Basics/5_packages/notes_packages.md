we have seen Go programs that have only one file with a main function and a couple of other functions. In real-world scenarios, this approach of writing all source code in a single file is not scalable. It becomes impossible to reuse and maintain code written this way. This is where packages are helpful.

Packages are used to organize Go source code for better reusability and readability.
Packages are a collection of Go sources files that reside in the same directory.
Packages provide code compartmentalization and hence it becomes easy to maintain Go projects.

==============================================================================

********\*\*********Go Modules**********\***********

Go Modules are needed to create custom packages.

A Go Module is nothing but a collection of Go packages.

The import path for the custom package we create is derived from the name of the go module.
All the other third-party packages(such as source code from github) along with their versions which our application uses will be managed by the go.mod file. This go.mod file is created when we create a new module.

==============================================================================

******\*******Creating a Go module********\*\*********

Run the command below to create a directory named learnpackage inside the current user’s Documents directory.

mkdir ~/Documents/learnpackage/

Make sure you are inside the directory learnpackage by typing cd ~/Documents/learnpackage/. Inside this directory run the following command to create a go module named learnpackage.

" go mod init learnpackage "

The above command will create a file named go.mod. The following will be the contents of the file.

module learnpackage

go 1.21.0

The line module learnpackage specifies that the module’s name is learnpackage. As we mentioned earlier, learnpackage will be the base path to import any package created inside this module. The line go 1.21.0 specifies that the files in this module use go version 1.21.0.

# ======================================================

****\*\*\*\*****Every executable Go application must contain the main function. This function is the entry point for execution. The main function should reside in the main package.****\*\*\*\*****

The line of code package main specifies that this file belongs to the main package. The import "packagename" statement is used to import an existing package. packagename.FunctionName() is the syntax to call a function in a package.

In line no. 3, we import the fmt package to use the Println function. The fmt is a standard package and is available as a part of the Go standard library. Then there is the main function which prints Simple interest calculation

=========================================================

To use a custom package we must import it first. The import path is the name of the go module concatenated by the directory where the package resides and the package name.

In our case the go module name is learnpackage and the package simpleinterest is in the simpleinterest folder directly under learnpackage

├── learnpackage
│ └── simpleinterest

So the line import "learnpackage/simpleinterest" will import the simpleinterest package.

=====================================================

In case we have a directory structure like this

learnpackage
│ └── finance
│ └── simpleinterest

then the import statement would be import "learnpackage/finance/simpleinterest"

============================================================

*********************Go Build**************************

Go tools like go build work in the context of the current directory.
  Till now we have been running go build from the directory ~/Documents/learnpackage/. If we try to run it from any other directory, it will fail.

Try coding into cd ~/Documents/ and then running go build learnpackage. It will fail with the following error.

package learnpackage is not in std (/usr/local/go/src/learnpackage)


=============================================================

go build has the ability to recursively search the parent directory for a go.mod file. Let’s check whether that works.

cd ~/Documents/learnpackage/simpleinterest/
The above command will take us to the simpleinterest directory. From that directory run

go build learnpackage
go build will successfully find a go.mod file in the parent directory learnpackage that has the module learnpackage defined and hence it works.

========================================================

It’s also possible to change the name of the output binary file using go build. Move to ~/Documents/learnpackage and type

go build -o fintechapp
The -o argument is used to specify the name of the output binary. In this case a binary file of name fintechapp will be created.

Run ./fintechapp and the binary will be run successfully.