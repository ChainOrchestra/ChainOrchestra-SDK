# [ChainOrchestra-SDK](https://github.com/ChainOrchestra/ChainOrchestra-SDK)

## Guest [registration](http://chainorchestra.net/#/4) example

### General design

The guest [registration](http://chainorchestra.net/#/4) live demo page is on [chainorchestra.net](http://chainorchestra.net).

### Sample session

```
cd examples/registration/js/
./guestList.js 
./deskOperation.js register "John Doe"
./guestStatus.js "John Doe"
./gateOperation.js checkin "John Doe"
./guestStatus.js "John Doe"
./gateOperation.js checkout "John Doe"
./guestStatus.js "John Doe"
./deskOperation.js unregister "John Doe"
./guestList.js 
```
