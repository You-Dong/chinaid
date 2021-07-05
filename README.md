## 注意

### 当您使用本项目时即代表您同意本项目以下协议:

**本项目仅用于测试目的，比如开发人员测试自己的关键校验规则是否正确；对于使用本项目产生的任何后果，使用者应当自行承担法律风险与责任，一切后果与本项目无关。**

## China ID

> chinaid 是一个用于生成中国各种信息的测试库，比如姓名、身份证号、地址、邮箱、银行卡号等。

本项目生成的测试数据尽量付合真实数据以模拟用户真实行为:

- 姓名: 使用常用的姓氏外加常见的名字，尽量使数据 "正常"
- 身份证号: 采用标准身份证规则生成(校验码有效)
- 手机号: 常用的手机号头部外加随机数字
- 银行卡号: 银行卡号采用正确的卡 bin 生成(LUHN 算法有效)
- 邮箱: 随机的前缀外加常用的域名后缀
- 地址: 省/城市信息使用真实数据，具体地址随机生成

## Benchmark

``` sh
➜ chinaid git:(master) ✗ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/mritd/chinaid
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkProvinceAndCity-16     62583489                16.55 ns/op            0 B/op          0 allocs/op
BenchmarkAddress-16              2399768               501.6 ns/op           223 B/op          6 allocs/op
BenchmarkLUHNProcess-16           905656              1384 ns/op             190 B/op         38 allocs/op
BenchmarkBankNo-16               1270488               936.0 ns/op           132 B/op         23 allocs/op
BenchmarkEmail-16                3394354               347.7 ns/op            56 B/op          5 allocs/op
BenchmarkIssueOrg-16            27343102                43.76 ns/op           11 B/op          0 allocs/op
BenchmarkValidPeriod-16          1329706               903.7 ns/op            40 B/op          3 allocs/op
BenchmarkIDNo-16                  914662              1307 ns/op             125 B/op         22 allocs/op
BenchmarkVerifyCode-16           3598204               328.2 ns/op            68 B/op         17 allocs/op
BenchmarkMobile-16               6248443               195.1 ns/op            32 B/op          3 allocs/op
BenchmarkName-16                16197261                70.51 ns/op           14 B/op          1 allocs/op
PASS
ok      github.com/mritd/chinaid        17.435s
```