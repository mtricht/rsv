# rsv

rsv is the newest, hottest file format. rsv stands for rsv-separated values.

The official specification is as follows:

1. Values are separated by a new line `\n`
2. Rows are separated by a comma
3. Escaping is not possible

## Example

```
name
age
gender
country,Michael
24
M
Belgium,Sven
28
M
Zimbabwe,Mara
24
F
Uganda,Peter
24
M
Belgium
```

View the full example file [here](example.rsv).

## Reference implementation

As this new file format can be hard to fathom, this repository also contains a reference implementation written in go. This reference implementation reads the `examples.rsv` and outputs the following to stdout:

```
➜ go run main.go
2018/05/04 15:50:28 --------------
2018/05/04 15:50:28 name: Michael
2018/05/04 15:50:28 age: 24
2018/05/04 15:50:28 gender: M
2018/05/04 15:50:28 country: Belgium
2018/05/04 15:50:28 --------------
2018/05/04 15:50:28 name: Sven
2018/05/04 15:50:28 age: 28
2018/05/04 15:50:28 gender: M
2018/05/04 15:50:28 country: Zimbabwe
2018/05/04 15:50:28 --------------
2018/05/04 15:50:28 name: Mara
2018/05/04 15:50:28 age: 24
2018/05/04 15:50:28 gender: F
2018/05/04 15:50:28 country: Uganda
```
