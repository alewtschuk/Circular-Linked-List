# Circular Linked List
This project is a circular linked list implementation translated from C into Go. It utilizes the unsafe package to hold generic types using unsafe.Pointer which is equivalent to C's void *. The tests have been translated as well as best as possible. There is no main as this is designed to implement and run the testing suite.

Steps to configure, build, run, and test the project.

## Building and Testing
Please test and build the project using the command:

```bash
make
```
This will auto run all tests in the testing package.

## Clean

```bash
make clean
```