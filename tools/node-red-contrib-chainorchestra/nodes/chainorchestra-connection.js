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
	var Connection = require('chainorchestra/chainorchestra.js').Connection;

	function ConnectionNode(config) {
		RED.nodes.createNode(this,config);
		this.proxy = new Connection(config.ip, config.port, config.user, config.password);
		this.params = this.context().flow.get('params') || {};
		this.context().flow.set('params', this.params);
		var node = this;
		this.on('input', function(msg) {
			if (msg.payload.enrollId)
				node.proxy.rpcObj.enrollId = msg.payload.enrollId;
			if (msg.payload.enrollSecret)
				node.proxy.rpcObj.enrollSecret = msg.payload.enrollSecret;
			node.proxy.login(
				function(obj) {
					node.params.secureContext = node.proxy.rpcObj.enrollId;
					msg.payload = obj;					
					node.send([msg,null]);
				},
				function(obj) {
					msg.payload = obj;
					node.send([null,msg]);
				});
		});
	}

	RED.nodes.registerType("connection",ConnectionNode);
};
