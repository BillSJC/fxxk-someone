# Fxxk someone

A Simple Proxy And Random UA Fxxker based on [zdaye](http://ip.zdaye.com/ShortProxy.html)

## Usage

Requirements

- zdaye`s API Key and Secret
- Go 1.11+

Consts

```go
const (
	API_KEY = "your api key"	//your API AppKey
	API_SECRET_MD5 = "your api secret (After MD5 16)"	//you API Secret (after MD5)
	TARGET_URL = "website u want to fuck"	//the website you want to fuck
	FREQUENCY = 100 // frequency (ms)
)
```

Build

```bash
go build
```