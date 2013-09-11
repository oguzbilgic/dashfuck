# Dashfuck [![Build Status](https://travis-ci.org/oguzbilgic/dashfuck.png)](https://travis-ci.org/oguzbilgic/dashfuck)

a little language that compiles into [brainfuck](http://en.wikipedia.org/wiki/Brainfuck)

## Installation

You can download the compiled `dashfuck` binary for your specific platform from the releases page. If you have the `go tools` installed you can compile the `dashfuck` with the command bellow:

```bash
$ go get github.com/oguzbilgic/dashfuck
``` 

## Usage

Dashfuck is almost like brainfuck language except the fact that it uses [dashes](http://en.wikipedia.org/wiki/Dash) to make more readable. 

Here is an example `hello.df` hello world program written in dashfuck:

```df
——————————[~———————~——————————~———~—<<<<-]~——–~—–———————––———–~——–<<———————————————–~–———–------–--------–~—–~–
```

Brainfuck | Dashfuck | Description
----------|----------|-------------
     +    |    —     | Em dash 
     .    |    –     | En dash 
     >    |    ~     | Swung dash 

To compile your dashfuck program into brainfuck, use the command bellow:

```bash
dashfuck hello.df
```

This will generate `main.bf` file which you can run via brainfuck vm.

## License

The MIT License (MIT)
