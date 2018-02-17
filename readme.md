Example for doing fizzbuzz using TDD in go

Some useful things

To run use

```
go run fizzbuzz.go
```

To test use
```
go test
```

To loop over array use

```
for index, element := range someSlice {
    // index is the index where we are
    // element is the element from someSlice for where we are
    // use _ in place of index or element if not using them
}
```

To convert string to int use
```
value, err := strconv.Atoi(stringThatMayBeANumber)
```

To convert int to number use
```
s := strconv.Itoa(myNumber)
```

To check divisible by a number use
```
if x%3 == 0 {
    // do something
}
```