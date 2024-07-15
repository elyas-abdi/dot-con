# .con

A new way to define and access your configurations dynamically based on custom patterns you define.

##### Here's an example `.con` File
```toml
# define default values at the top of the file
title = an example dot con file
key = default value for this key

# you can define wildcards explicitly, 
[stage: prod] [region: *]
key = value for prod stage regardless of what region is set

# or you can omit the pattern element/s to achive the same effect
[stage: prod]
key = value for prod stage regardless of what region is set

# you can also define multiple values for your pattern elements
[stage: alpha | beta | gamma] [region: us-east-2]
key = shared value for test stages (alpha, beta, or gamma) in us-east-2
another_key = another value

# use of string arrays, int, floats, chars are supported in addition to string types 
[region: us-east-2] [stage: prod-na] # <- you can also define the pattern in any order
some_int = 987
some_float = 3.141592653589793238462643383279502884197
some_array_of_strings = [ele1, ele2, ele3, ele4]
some_char = c

# you can add as many pattern elements as you wish
[region: us-east-2] [stage: prod-na] [experimental: true] [some other pattern element: pattern element one] [another pattern element: pattern element two]
some really specific key = some really specifc value for this scenario
```
##### Here's how to use the client
```go
package main

import `github.com/elyas-abdi/dot-con`

func main() {
    con, err := config.New().Load()
    if err != nil {
        print(err)
    }
	
    value := con.String("key").Access()
}
```


## ⬇️ Installation 

Run this command in your project's directory

```shell
go get github.com/elyas-abdi/dot-con
```
> ⚠️ **CAUTION**: Dot Con is currently in experimental development. Use in production at your own discretion 

## 🦮️ Usage Guide

##### Step 1 - Create a directory to store your configuration files.
You can name your directory whatever you prefer. However, the Dot Con client will default to looking for a folder named `con` in your root project directory when searching for `.con` files, unless you specify a custom directory path when initializing the client.

##### Step 2 - Create a .con file to define your config variables in
Inside the directory you created in step 1, create a new file that ends in `.con`. This is where you will define your config variables and access patterns.

If you are sticking with the default layout that Dot Con uses, then your project structure will look something like this:
```text
├── your-project-directory
│   ├── con
│   │   ├── {any file name}.con
│   ├── go.mod
│   ├── go.sum
│   ├── ...
```
You can define and use multiple `.con` files in the same directory, and the client will concurrently parse all of them at once.

##### Step 3 - Define your config variables
todo... 
## 🤓 Under the Hood~~
todo... 