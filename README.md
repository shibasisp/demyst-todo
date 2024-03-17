# demyst-todo


## Table of Contents

- [demyst-todo](#demyst-todo)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Usage](#usage)

## Introduction

demyst-todo is a cli tool that consumes todos from API in most performant way and outputs the title and status in the TTY.
We can pass the number of todos to fetch and the pattern(like even, odd etc).

## Prerequisites
Ensure you have these tools installed in your system to run this.
1. Docker
or
2. Golang

## Usage
```
❯ ./demyst-todo status help
NAME:
   demyst-todo status - Returns the status of the TODO

USAGE:
   demyst-todo status [command options] [arguments...]

OPTIONS:
   --limit value, -l value        The number of todos to fetch (default: 20)
   --pattern value, -p value      The pattern to filter the todos. Available Pattern: even, odd, all (default: "all")
   --input value, -i value        The input source of the todos. Available values: api, file (default: "api")
   --location value, --loc value  The location to fetch the todos from. It can either be file location or API url depending on the input (default: "https://jsonplaceholder.typicode.com/todos")
   --help, -h                     show help
```
**Docker**
1. Build the image from the project
```
docker build -t demyst .                                                                                                                                                                                                                       
```
2. Run the project based on various flags shown above.
```
docker run --rm -e "FLAGS=-l 3" demyst
```
**Golang**

1. Install the dependencies in your system.
```
go get
```
2. Run the project based on various flags.
```
make run FLAGS="-limit 3 -pattern even"                                                                                                                                                                                                             
```
