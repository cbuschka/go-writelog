# go-writelog

### Command line tool that reads lines from stdin and writes them to a target file with timestamp prefix

## Usage

```
[conni@dudu go-writelog]$ echo -ne "line1\nline2" | writelog test.log
[conni@dudu go-writelog]$ cat test.log 
2022-01-16T22:11:46 line1
2022-01-16T22:11:46 line2
```

## License

Copyright (c) 2022 by [Cornelius Buschka](https://github.com/cbuschka).

[MIT](./license.txt)
