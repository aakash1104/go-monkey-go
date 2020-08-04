<p align="center">
    <img src="resources/monkey-drawing.png" alt="Monkey Programming Language"
    height="25%" width="25%">
</p>

<h2 align="center">
    :monkey: MonkeyLang Interpreter :monkey:
    <br>
    <img src="https://img.shields.io/badge/status-under%20development-yellow">
</h2>

`Go-Monkey-Go` is an interpreter for [monkeylang](https://monkeylang.org/), a
simple programming language created by 
[Thorsten Ball](https://thorstenball.com/) written in Go. This is adapted from
Ball's book [Writing an Interpreter in Go](https://interpreterbook.com).

###  Understanding MonkeyLang Syntax ###

Monkey has a very simple syntax. Monkey supports primitive data types such as
integers, strings, booleans. Variable declarations are done like this:

```js
let num = 10;
let b = num + 50;
let str = "Hello";
let flag = true;
```

Monkey also supports user defined functions:

```js
let sum = fn(a, b) {
    x + y;
};
```
There's also support for conditional statements:

```js
if (a <= b) {
    return true;
} else {
    return false;
}
```

### Motivation :thinking: ### 

After reading this [interview](https://evrone.com/rob-pike-interview) of Rob 
Pike, I was extremely curious about Go and it's advantages in developing and 
deploying software. 

One of the best ways to get exposed to a programming language is to build 
something from scratch. Compilers and Interpreters are one of the most 
beautiful pieces of art in Computer Science. So here's my attempt of building an
interpreter in Go. My goal is to get thinking about Go and Monkeylang and add 
other features to this programming language (I need to think more about what I 
want to add though).

_Note_: This repository will be constantly updated as I read the book and 
complete the exercises. Stay tuned! :nerd_face:

