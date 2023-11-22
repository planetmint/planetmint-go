# Library for RPC requests to Planetmint

## How to use it

In the example below we use the account `addr0` for which we have the private key in our keyring.
We configure the address prefix andd change the default RPC endpoint to a remote one.
The only keyring backend currently supported is the test backend under `keyring-test`.
After that we construct three messages to send `10plmnt` each to three addresses `addr1`, `addr2` and `addr3`.
We then build and sign the transaction and eventually send this transaction via RPC.
For debugging purposes we print the transaction that we send as JSON.

```
package main

import (
        "fmt"
        "log"

        sdk "github.com/cosmos/cosmos-sdk/types"
        banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
        "github.com/planetmint/planetmint-go/app"
        "github.com/planetmint/planetmint-go/lib"
        "golang.org/x/net/context"
)

func main() {
        encodingConfig := app.MakeEncodingConfig()

        libConfig := lib.GetConfig()
        libConfig.SetEncodingConfig(encodingConfig)
        libConfig.SetAPIEndpoint("https://testnet-api.rddl.io")
        libConfig.SetRPCEndpoint("https://testnet-rpc.rddl.io:443")

        addr0 := sdk.MustAccAddressFromBech32("plmnt168z8fyyzap0nw75d4atv9ucr2ye60d57dzlzaf")
        addr1 := sdk.MustAccAddressFromBech32("plmnt1vklujvmr9hsk9zwpquk4waecr2u5vcyjd8vgm8")
        addr2 := sdk.MustAccAddressFromBech32("plmnt1pwquxvqmmdry4gdel4g4rz0js7jy65453h92g7")
        addr3 := sdk.MustAccAddressFromBech32("plmnt1dyuhg8ldu3d6nvhrvzzemtc3893dys9v9lvdty")

        coin := sdk.NewCoins(sdk.NewInt64Coin("plmnt", 10))
        msg1 := banktypes.NewMsgSend(addr0, addr1, coin)
        msg2 := banktypes.NewMsgSend(addr0, addr2, coin)
        msg3 := banktypes.NewMsgSend(addr0, addr3, coin)

        ctx := context.Background()
        txJSON, err := lib.BuildUnsignedTx(ctx, addr0, msg1, msg2, msg3)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(txJSON)

        _, err = lib.BroadcastTx(ctx, addr0, msg1, msg2, msg3)
        if err != nil {
                log.Fatal(err)
        }
}
```

Sample output:
```
$ go run main.go|jq
{
  "body": {
    "messages": [
      {
        "@type": "/cosmos.bank.v1beta1.MsgSend",
        "from_address": "plmnt168z8fyyzap0nw75d4atv9ucr2ye60d57dzlzaf",
        "to_address": "plmnt1vklujvmr9hsk9zwpquk4waecr2u5vcyjd8vgm8",
        "amount": [
          {
            "denom": "plmnt",
            "amount": "10"
          }
        ]
      },
      {
        "@type": "/cosmos.bank.v1beta1.MsgSend",
        "from_address": "plmnt168z8fyyzap0nw75d4atv9ucr2ye60d57dzlzaf",
        "to_address": "plmnt1pwquxvqmmdry4gdel4g4rz0js7jy65453h92g7",
        "amount": [
          {
            "denom": "plmnt",
            "amount": "10"
          }
        ]
      },
      {
        "@type": "/cosmos.bank.v1beta1.MsgSend",
        "from_address": "plmnt168z8fyyzap0nw75d4atv9ucr2ye60d57dzlzaf",
        "to_address": "plmnt1dyuhg8ldu3d6nvhrvzzemtc3893dys9v9lvdty",
        "amount": [
          {
            "denom": "plmnt",
            "amount": "10"
          }
        ]
      }
    ],
    "memo": "",
    "timeout_height": "0",
    "extension_options": [],
    "non_critical_extension_options": []
  },
  "auth_info": {
    "signer_infos": [],
    "fee": {
      "amount": [
        {
          "denom": "plmnt",
          "amount": "1"
        }
      ],
      "gas_limit": "200000",
      "payer": "",
      "granter": ""
    },
    "tip": null
  },
  "signatures": []
}
```
