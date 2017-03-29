# [ChainOrchestra-SDK](https://github.com/ChainOrchestra/ChainOrchestra-SDK)

## Value [exchange](http://chainorchestra.net/#/5) example

### General design

The value [exchange](http://chainorchestra.net/#/5) live demo page is on [chainorchestra.net](http://chainorchestra.net).

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
