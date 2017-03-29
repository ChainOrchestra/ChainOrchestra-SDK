#!/usr/bin/env node

function usage() {
	console.log("Usage:\n\t" + process.argv[1].split("/").pop() + " <fromAccount> <toAccount> <amount>")
}

if (process.argv.length != 5) {
	usage();
	process.exit(1);
}

var log = console.log;

var Connection = require("chainorchestra/chainorchestra.js").Connection
var Query = require('chainorchestra/chainorchestra.js').Query
var Transaction = require("chainorchestra/chainorchestra.js").Transaction

var ip = "51.15.43.252";
var port = "7050";

var user = "Cash01";
var chaincodeID = "3a4a643bf2591f5667ea3c6b268760360becab578620bfe5fea8b87a54e5b62fe159992552a3a88be23ae685dac755c5850796e722198369de6d35bd37991833";

var transaction = new Transaction(ip, port);
transaction.setSecureContext(user);
transaction.setChaincodeID("name", chaincodeID);

function success(result) {
  log("Transfer of " + this.amountValue + " from account " + this.fromAccountName + " to account " + this.toAccountName + " success");
  log("Transaction id: " + result.result.message);
}

function failure(result) {
  log("Transfer failure from account " + this.fromAccountName + " to account " + this.fromAccountName);
  log(JSON.stringify(result));
}

var fromAccount = process.argv[2];
var toAccount = process.argv[3];
var amount = process.argv[4];

var context = transaction.invoke("transfer", [fromAccount, toAccount, amount], success, failure);
context.fromAccountName = fromAccount;
context.toAccountName = toAccount;
context.amountValue = amount;
