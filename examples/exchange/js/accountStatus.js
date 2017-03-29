#!/usr/bin/env node

function usage() {
	console.log("Usage:\n\t" + process.argv[1].split("/").pop() + " <account>")
}

if (process.argv.length != 3) {
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
var chaincodeID = "3a4a643bf2591f5667ea3c6b268760360becab578620bfe5fea8b87a54e5b62fe159992552a3a88be23ae685dac755c5850796e722198369de6d35bd37991833";

var query = new Query(ip, port);
query.setSecureContext(user);
query.setChaincodeID("name", chaincodeID);

function success(result) {
  log("Account " + this.accountName + " status:");
  log(JSON.stringify(result));
}

function failure(result) {
  log("Failed retrieving account " + this.accountName + " status.");
  log(JSON.stringify(result));
}

var account = process.argv[2];

var context = query.query("query", [account], success, failure);
context.accountName = account;
