# [ChainOrchestra-SDK](https://github.com/ChainOrchestra/ChainOrchestra-SDK)

## ChainOrchestra simplified Hyperledger REST client API for Node.js

### Installing the chainorchestra API for Node.js

You need to install [Node.js](https://nodejs.org/) / npm first.

```
sudo apt-get install -y nodejs nodejs-legacy npm
```

Then install the chainorchestra API for Node.js with npm

```
cd lib/js
sudo npm install -g
```

Once chainorchestra API for Node.js is installed, you can also install the 
[ChainOrchestra blockchain nodes for Node-RED](../../tools/node-red-contrib-chainorchestra)


### [chainorchestra.js](https://chainorchestra.github.io/ChainOrchestra-SDK/chainorchestra.js.html) library usage

Import the [API object classes](https://chainorchestra.github.io/ChainOrchestra-SDK/index.html) as needed.

```
var Connection = require("chainorchestra/chainorchestra.js").Connection
var Query = require('chainorchestra/chainorchestra.js').Query
var Transaction = require("chainorchestra/chainorchestra.js").Transaction
```

The code snippets from the web app samples can be used unchanged.

  * [Connection sample](http://chainorchestra.net/ChainOrchestra-SDK/sampleConnection.html)
  * [Query sample](http://chainorchestra.net/ChainOrchestra-SDK/sampleQuery.html)
  * [Transaction sample](http://chainorchestra.net/ChainOrchestra-SDK/sampleTransaction.html)


