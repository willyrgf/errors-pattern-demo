# errors-pattern-demo

```sh
$ go run ./main.go
a standard error with exit code: something goes wrong in pkg1: open foo.bar: no such file or directory
something goes wrong in pkg1: open foo.bar: no such file or directory
true
true
true
os.Exit(): 1

a specific error with exit code: something goes wrong in pkg2: open foo.bar: no such file or directory
something goes wrong in pkg2: open foo.bar: no such file or directory
true
true
true
os.Exit(): 127
```