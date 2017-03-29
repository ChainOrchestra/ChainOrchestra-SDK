#!/usr/bin/env node

function usage() {
	console.log("Usage:\n\t" + process.argv[1].split("/").pop())
}

if (process.argv.length != 2) {
	usage();
	process.exit(1);
}

var log = console.log;

var Connection = require("chainorchestra/chainorchestra.js").Connection
var Query = require('chainorchestra/chainorchestra.js').Query
var Transaction = require("chainorchestra/chainorchestra.js").Transaction

var ip = "51.15.43.252";
var port = "7050";

var user = "DemoUser";
var chaincodeID = "f519d9cc955c612fb573332166e08fa8ea3682facbd234545d7dcdfe1554e12cd11673fd90bf4716467ad6ed2e90a2f8a6da7139b24eb19f59be1f14f8fc4d67";

var query = new Query(ip, port);
query.setSecureContext(user);
query.setChaincodeID("name", chaincodeID);

function success(result) {
  var guests = result.result.message.split("\n");
  guests.pop();
  log(guests);
}

function failure(result) {
  log("Failed retrieving guest list.");
  log(JSON.stringify(result));
}

var context = query.query("list", [], success, failure);
