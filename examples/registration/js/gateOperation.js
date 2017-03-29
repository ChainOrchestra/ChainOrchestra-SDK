#!/usr/bin/env node

function usage() {
	console.log("Usage:\n\t" + process.argv[1].split("/").pop() + " <checkin|checkout> <guest>")
}

if (process.argv.length != 4) {
	usage();
	process.exit(1);
}

var log = console.log;

var Connection = require("chainorchestra/chainorchestra.js").Connection
var Query = require('chainorchestra/chainorchestra.js').Query
var Transaction = require("chainorchestra/chainorchestra.js").Transaction

var ip = "51.15.43.252";
var port = "7050";

var user = "Guard1";
var chaincodeID = "f519d9cc955c612fb573332166e08fa8ea3682facbd234545d7dcdfe1554e12cd11673fd90bf4716467ad6ed2e90a2f8a6da7139b24eb19f59be1f14f8fc4d67";

var transaction = new Transaction(ip, port);
transaction.setSecureContext(user);
transaction.setChaincodeID("name", chaincodeID);

function success(result) {
  log(context.gateOp + " success for guest " + this.guestName);
  log("Transaction id: " + result.result.message);
}

function failure(result) {
  log(context.gateOp + " failure for guest " + this.guestName);
  log(JSON.stringify(result));
}

var operation = process.argv[2];
var guest = process.argv[3];

var context = transaction.invoke(operation, [guest], success, failure);
context.gateOp = operation;
context.guestName = guest;
