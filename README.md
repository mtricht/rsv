# rsv

rsv is the newest, hottest file format. rsv stands for rsv-separated values.

The official specification is as follows:

1. Values are separated by a new line `\n`
2. Rows are separated by a comma
3. Escaping is not possible

## Examples

An example .rsv file can be found [here](example.rsv).

This repository also contains a small go implementation of this awesome new file format. The program outputs:

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
