
# Wrap Up

- Declaring structs to create your own data types which lets you bundle related
data together and make the intent of your code clearer
- Declaring interfaces so you can define functions that can be used by different
types
- Adding methods so you can add functionality to your data types and so you can
implement interfaces
- Table driven tests (TDT) makes you assertions clearer and you test suites
easier to extend & maintain

Interfaces are a great tool for hiding complexity away from other parts of the
system. like in our test code helper code did not need to know the exact shape
it was asserting on, only how to "ask" for its area.
