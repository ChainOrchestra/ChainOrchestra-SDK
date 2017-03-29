# [ChainOrchestra-SDK](https://github.com/ChainOrchestra/ChainOrchestra-SDK)

This SDK provides tools and APIs to access the ChainOrchestra blockchains based on 
[Hyperledger v0.6](http://hyperledger-fabric.readthedocs.io/en/v0.6/API/CoreAPI.html)

The software in this repository is distributed in open-source under the 
[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0)


## ChainOrchestra simplified Hyperledger REST client API

### [API documentation](https://chainorchestra.github.io/ChainOrchestra-SDK/index.html)

The [chainorchestra.js API](https://chainorchestra.github.io/ChainOrchestra-SDK/index.html) defines 
a basic set of objects to handle Hyperledger peer interactions over the REST API.

### Using the ChainOrchestra client API in a Web app

Just include the [chainorchestra.js](https://chainorchestra.github.io/ChainOrchestra-SDK/chainorchestra.js.html) lib in a web page
to connect to Hyperledger using the REST API.

```
<script src="chainorchestra.js"></script>
```

The following live examples from the [registration](http://chainorchestra.net/#/4) application on 
[chainorchestra.net](http://chainorchestra.net) use the chainorchestra.js objects.

  * [Connection live example](http://chainorchestra.net/ChainOrchestra-SDK/sampleConnection.html)
  * [Query live example](http://chainorchestra.net/ChainOrchestra-SDK/sampleQuery.html)
  * [Transaction live example](http://chainorchestra.net/ChainOrchestra-SDK/sampleTransaction.html)

### Using the [ChainOrchestra client API with Node.js](lib/js/)

Install the [chainorchestra.js](https://chainorchestra.github.io/ChainOrchestra-SDK/chainorchestra.js.html) lib 
as a Node.js package to run scripts and use it as a command line app.


## ChainOrchestra chaincode and server side APIs

Chaincode APIs provide services shared by all user chaincode applications deployed on a ChainOrchestra-managed, mutualized blockchain.
It can be libraries to include or interfaces to implement in order to access a specific service.

### Chaincode [permissions module](lib/chaincode/perms)

The [perms.go](lib/chaincode/perms/perms.go) permissions module is a library that chaincode can include in order to 
manage consistent, prioritized permissions across different applications.


## ChainOrchestra blockchain demos

  * Live demo blockchain on [chainorchestra.net](http://chainorchestra.net)
  * Guest [registration](http://chainorchestra.net/#/4) application
  * Value [exchange](http://chainorchestra.net/#/5) application


## About ChainOrchestra

[ChainOrchestra](http://chainorchestra.com) conceives and develops demonstrators, 
proof-of-concept and pilots on its own permissioned blockchains.

