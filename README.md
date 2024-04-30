# morpho-liquidator-bot

## Contract
`cd contract && forge install`

## Config
```yaml
http_endpoint: "" # http endpoint for the ethereum node
ws_endpoint: "" # ws endpoint for the ethereum node
bot_address: "" # address of the bot
private_key: "" # private key of the eoa account
one_inch_api_key: "" # api key for 1inch (https://portal.1inch.dev/)
markets: # list of markets to monitor (https://app.morpho.org/borrow)
  - "0xc581c5f70bd1afa283eed57d1418c6432cbff1d862f94eaf58fdd4e46afbb67f" #usde-dai
  - "0xc54d7acf14de29e0e5527cabd7a576506870346a78a11a6762e2cca66322ec41" # wsteth-weth
  - "0x698fe98247a40c5771537b5786b2f3f9d78eb487b4ce4d75533cd0e94d88a115" # weeth-weth
```

## Run
`go build . && ./morpho-liquidator-bot`
