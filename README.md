# wecho

`echo` command for Windows cmd.exe

echo command on cmd.exe doesn't work with UTF-8 strings. When call DOS batch file like below, we want the redirected file is encoded in UTF-8. But when not redirected, we want to look correctly on cmd.exe .

```dosbatch
wecho こんにちわ世界
```

```
C:\>foo.bat
こんにちわ世界

C:\>foo.bat > output
```

## Installation

```
$ go get github.com/mattn/wecho
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
