/******************************************************************************
 * Copyright 2016 ChainOrchestra.
 * 
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * 
 * http://www.apache.org/licenses/LICENSE-2.0
 * 
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *****************************************************************************/


module.exports = function(RED) {
	var Transaction = require('chainorchestra/chainorchestra.js').Transaction;

	function TransactionNode(config) {
		RED.nodes.createNode(this,config);
		this.proxy = new Transaction(config.ip, config.port);
		this.proxy.setSecureContext(config.user);
		this.proxy.setChaincodeID("name", config.chaincodeid);
		this.params = this.context().flow.get('params') || {};
		this.context().flow.set('params', this.params);
		var node = this;
		this.on('input', function(msg) {
			for (var key in node.params) {
				this.proxy.rpcObj.params[key] = node.params[key];
			}
			for (var key in msg.payload) {
				this.proxy.rpcObj.params[key] = msg.payload[key];
			}
			node.proxy.invoke(
				node.proxy.rpcObj.params.ctorMsg.function,
				node.proxy.rpcObj.params.ctorMsg.args,
				function(obj) {
					msg.payload = obj;
					node.send([msg,null]);
				},
				function(obj) {
					msg.payload = obj;
					node.send([null,msg]);
				});
		});
	}

	RED.nodes.registerType("transaction",TransactionNode);
};
