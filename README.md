# Demo-Crypto CLI

__CLI app for demo trading/investing in crypto written in Go__

## Installation

Add /build package to PATH

Add environment variable API_KEY  

Compile code with command:
```powershell
make compile
```

init CLI storage with command:

```powershell
crypto init
```

## Commands:

* ```crypto profile``` - prints your user profile (name, balance, wallet)

* ```crypto price``` - prints current crypto prices

* ```crypto buy``` - buys crypto (use flags -id, -usd)

* ```crypto sell``` - sell crypto (use flags -id, -usd)

* soon...
