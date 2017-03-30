# [ChainOrchestra-SDK](https://github.com/ChainOrchestra/ChainOrchestra-SDK)

## Value [exchange](http://chainorchestra.net/#/5) example

### General design

This sample digital currency application opens accounts, manages deposits, withdrawals and transfers.
The balance needs to be 0 before closing an account.

In the blockchain DB, the account name is the key and the value is the account balance.

The value [exchange](http://chainorchestra.net/#/5) live demo page is on [chainorchestra.net](http://chainorchestra.net).

#### Desk operations

Blockchain users with [BankClerk permissions](../../lib/chaincode/perms) manage the open / close operations.

```
---- ---- ---- ---- ---- ----+---- ---- ---- ----+---- ---- ---- ---- ---- ---- 
     Initial balance         |      Operation    |         Next balance
---- ---- ---- ---- ---- ----+---- ---- ---- ----+---- ---- ---- ---- ---- ---- 
                             |                   |
     <undefined>             |        open       |             0
                             |                   |
          0                  |       close       |         <undefined> 
                             |                   |
---- ---- ---- ---- ---- ----+---- ---- ---- ----+---- ---- ---- ---- ---- ---- 
 
```

#### Teller operations

Blockchain users with [Cashier permissions](../../lib/chaincode/perms) manage the deposit / withdrawal / transfer operations.

```
---- ---- ---- ---- ---- ----+---- ---- ---- ----+---- ---- ---- ---- ---- ---- 
     Initial balance         |      Operation    |         Next balance
---- ---- ---- ---- ---- ----+---- ---- ---- ----+---- ---- ---- ---- ---- ---- 
                             |                   |
           N                 |     deposit x     |          N + x
                             |                   |
           N                 |    withdrawal x   |          N - x
                             |                   |
   fromAccount: N            |     transfer x    |          N - x
    toAccount: M             |                   |          M + x
                             |                   |
---- ---- ---- ---- ---- ----+---- ---- ---- ----+---- ---- ---- ---- ---- ---- 
 
```


### Sample session

```
cd examples/exchange/js/
./accountList.js 
./deskOperation.js open "John Doe"
./tellerOperation.js deposit "John Doe" 100
./accountStatus.js "John Doe"
./deskOperation.js open "Joe Black"
./transferOperation.js "John Doe" "Joe Black" 10
./accountStatus.js "Joe Black"
./tellerOperation.js withdrawal "John Doe" 90
./tellerOperation.js withdrawal "Joe Black" 10
./deskOperation.js close "John Doe"
./deskOperation.js close "Joe Black"
./accountList.js 
```
