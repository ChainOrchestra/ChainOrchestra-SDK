# [ChainOrchestra-SDK](https://github.com/ChainOrchestra/ChainOrchestra-SDK)

## ChainOrchestra blockchain nodes for [Node-RED](https://nodered.org/)

Node-RED is a programming tool built on Node.js with a browser-based flow diagram editor.
The blockchain nodes use the [ChainOrchestra API for Node.js](../../lib/js) to 
expose the API objects as flow boxes.

### Installing the blockchain nodes

Before using the blockchain nodes for Node-RED, you need to install the following components:

  * [Node.js](https://nodejs.org/)
  * [Node-RED](https://nodered.org/)

```
sudo apt-get install -y nodejs nodejs-legacy npm
sudo npm install -g node-red
```

Then install the [ChainOrchestra API for Node.js](../../lib/js) directly from this SDK,
and install the chainorchestra blockchain nodes with npm.

```
cd  tools/node-red-contrib-chainorchestra
sudo npm install -g
```

### Using the blockchain nodes

There are 4 types of blockchain nodes, reflecting the [ChainOrchestra API objects](https://chainorchestra.github.io/ChainOrchestra-SDK/index.html)

  * The **Connection node** makes a login connection.
  * The **Chaincode node** deploys chaincode.
  * The **Query node** runs a query.
  * The **Transaction node** executes a transaction.
