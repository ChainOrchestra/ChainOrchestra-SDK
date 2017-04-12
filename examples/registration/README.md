# [ChainOrchestra-SDK](https://github.com/ChainOrchestra/ChainOrchestra-SDK)

## Guest [registration](http://chainorchestra.net/#/4) example

### General design

This live event / hospitality sample application registers guests, online or at a desk, then let them checkin at a gate.
The guest needs to checkout before unregistering.

In the blockchain DB, the user name is the key, the value of wich is either:

  * ```registered``` if the user is present in the DB but has not physically passed the gate
  * ```checked in``` if the user entered the event

The guest [registration](http://chainorchestra.net/#/4) live demo page is on [chainorchestra.net](http://chainorchestra.net).

#### Desk operations

Blockchain users with [DeskClerk permissions](../../lib/chaincode/perms) manage the register / unregister operations.

```
---- ---- ---- ---- ---- ----+---- ---- ---- ----+---- ---- ---- ---- ---- ---- 
     Initial state           |      Operation    |         Next state
---- ---- ---- ---- ---- ----+---- ---- ---- ----+---- ---- ---- ---- ---- ---- 
                             |                   |
     <undefined>             |      register     |         registered
                             |                   |
     registered              |     unregister    |         <undefined> 
                             |                   |
---- ---- ---- ---- ---- ----+---- ---- ---- ----+---- ---- ---- ---- ---- ---- 
 
```

#### Gate operations

Blockchain users with [SecurityGuard permissions](../../lib/chaincode/perms) manage the checkin / checkout operations.

```
---- ---- ---- ---- ---- ----+---- ---- ---- ----+---- ---- ---- ---- ---- ---- 
     Initial state           |      Operation    |         Next state
---- ---- ---- ---- ---- ----+---- ---- ---- ----+---- ---- ---- ---- ---- ---- 
                             |                   |
     registered              |      checkin      |         checked in 
                             |                   |
     checked in              |      checkout     |         registered
                             |                   |
---- ---- ---- ---- ---- ----+---- ---- ---- ----+---- ---- ---- ---- ---- ---- 
 
```


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

### Sample Node-RED gateway flows

Copy the json samples in the [flows](flows) folder to the ```~/.node-red/lib/flows/``` folder on your Node-RED gateway,
then import them on the Node-RED web interface using the **Import -> Library** menu.

The flow samples are documented with the [ChainOrchestra flow nodes](https://chainorchestra.github.io/ChainOrchestra-SDK/flowNodes.html).

