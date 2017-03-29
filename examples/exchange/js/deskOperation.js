#!/usr/bin/env node

function usage() {
	console.log("Usage:\n\t" + process.argv[1].split("/").pop() + " <open|close> <account>")
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

var user = "BankC1";
var chaincodeID = "3a4a643bf2591f5667ea3c6b268760360becab578620bfe5fea8b87a54e5b62fe159992552a3a88be23ae685dac755c5850796e722198369de6d35bd37991833";

var transaction = new Transaction(ip, port);
transaction.setSecureContext(user);
transaction.setChaincodeID("name", chaincodeID);

function success(result) {
  log(context.deskOp + " success for account " + this.accountName);
  log("Transaction id: " + result.result.message);
}

function failure(result) {
  log(context.deskOp + " failure for account " + this.accountName);
  log(JSON.stringify(result));
}

var operation = process.argv[2];
var account = process.argv[3];

var context = transaction.invoke(operation, [account], success, failure);
context.deskOp = operation;
context.accountName = account;
