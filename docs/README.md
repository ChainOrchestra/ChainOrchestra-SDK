# [ChainOrchestra-SDK](https://github.com/ChainOrchestra/ChainOrchestra-SDK)

## ChainOrchestra simplified Hyperledger REST client [API](https://chainorchestra.github.io/ChainOrchestra-SDK/index.html)

Samples and documentation for the main API classes 

### PeerProxy

The ChainOrchestra peer proxy manages all asynchronous calls to a Hyperledger peer. 
It is the base class for other REST API helper objects described here. 

  * [PeerProxy](PeerProxy.html) class API documentation

### Connection

Connects a user to a network peer on the blockchain. 
As of Hyperledger v0.6 member services, this is done once and for all on a given peer.

  * [Connection](Connection.html) class API documentation
  * Sample [Connection](http://chainorchestra.net/ChainOrchestra-SDK/sampleConnection.html) code live example on the Registration and gate access application.
 
### Query

Connects to a single peer to query the blockchain.

  * [Query](Query.html) class API documentation
  * Sample [Query](http://chainorchestra.net/ChainOrchestra-SDK/sampleQuery.html) code live example on the Registration and gate access application.
 
### Transaction

Runs a transaction on the blockchain, verify it is validated by peer consensus.

  * [Transaction](Transaction.html) class API documentation
  * Sample [Transaction](http://chainorchestra.net/ChainOrchestra-SDK/sampleTransaction.html) code live example on the Registration and gate access application.
 

## Hyperledger Fabric Client / REST compatibility

The ChainOrchestra API provides a minimal compatibility layer for using the HFC objects over the Hyperledger REST API.

### Chain

  * [Chain](Chain.html) API wrapper over the HFC Chain class.
  
### Member

  * [Member](Member.html) API wrapper over the HFC Member class.
