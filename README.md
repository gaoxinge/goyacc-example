# goyacc example

## install goyacc

Following [goyacc](https://pkg.go.dev/modernc.org/goyacc) and [go install](https://stackoverflow.com/a/68809471/5772365), we can install goyacc in golang 1.18 by command:

```
$ go install modernc.org/goyacc@latest
```

## content

- [example1](./example1): plain calculator
- [example2](./example2): ast calculator
- [example3](./example3): tidb mysql parser example

## reference

- [TiDB 源码阅读系列文章（五）TiDB SQL Parser 的实现](https://pingcap.com/zh/blog/tidb-source-code-reading-5)
- [slrtbtfs/goyacc-tutorial](https://github.com/slrtbtfs/goyacc-tutorial)
- [tkrs/parser.go.y](https://gist.github.com/tkrs/23ab96cf84284ec6848c1c0a16611e21)
- [pingcap/parser](https://github.com/pingcap/parser/blob/master/docs/quickstart.md)
