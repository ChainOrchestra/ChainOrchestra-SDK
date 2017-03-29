/*
 Copyright ChainOrchestra All Rights Reserved.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"chainorchestra.com/perms"
)

type ExchangeChaincode struct {
}


func (t *ExchangeChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error

	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	permissions, err := perms.CheckPermissions(stub)
	if err != nil {
		return nil, err
	}
	
	if (0 == permissions & perms.SystemOperator) {
		return nil, errors.New("Need System Operator rights to deploy")
	}
	
	return nil, err
}


func (t *ExchangeChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	var acc1, acc2 string
	var val, val1, val2 int

	permissions, err := perms.CheckPermissions(stub)
	if err != nil {
		return nil, err
	}
	
	if function == "open" || function == "close" {
		if len(args) != 1 {
			return nil, errors.New("Expected argument: Account")
		}
		if (0 == permissions & perms.BankClerk) {
			return nil, errors.New("Need Bank Clerk rights to open or close account")
		}
		acc1 = args[0]
		rawBytes, err := stub.GetState(acc1);
		if rawBytes == nil {
			if function == "close" {
				return nil, errors.New("Cannot close unknown account " + acc1)
			}
			err = stub.PutState(acc1, []byte(strconv.Itoa(0)))
		} else {
			if function == "open" {
				return nil, errors.New("Cannot open already existing account " + acc1)
			}
			val1, _ = strconv.Atoi(string(rawBytes))
			if val1 != 0 {
				return nil, errors.New("Cannot close account " + acc1 + ", balance is not zero")
			}
			err = stub.DelState(acc1)
		}
		return nil, err
	}
	
	if function == "deposit" || function == "withdrawal" {
		if len(args) != 2 {
			return nil, errors.New("Expected arguments: Account, Value")
		}
		if (0 == permissions & perms.Cashier) {
			return nil, errors.New("Need Cashier rights for exchange operations")
		}
		acc1 = args[0]
		rawBytes, err := stub.GetState(acc1);
		if rawBytes == nil {
			return nil, errors.New("Cannot access unknown account " + acc1)
		}
		val1, _ = strconv.Atoi(string(rawBytes))
		val, err = strconv.Atoi(args[1])
		if err != nil {
			return nil, errors.New("Expected integer value, got " + args[1])
		}
		if function == "deposit" {
			val1 = val1 + val
		} else {
			val1 = val1 - val
		}
		err = stub.PutState(acc1, []byte(strconv.Itoa(val1)))
		return nil, err
	}
	
	if function == "transfer" {
		if len(args) != 3 {
			return nil, errors.New("Expected arguments: FromAccount, ToAccount, Value")
		}
		if (0 == permissions & perms.Cashier) {
			return nil, errors.New("Need Cashier rights for exchange operations")
		}
		acc1 = args[0]
		acc2 = args[1]
		if acc1 == acc2 {
			return nil, errors.New("Cannot transfer to the same account " + acc1)
		}
		rawBytes, err := stub.GetState(acc1);
		if rawBytes == nil {
			return nil, errors.New("Cannot access unknown account " + acc1)
		}
		val1, _ = strconv.Atoi(string(rawBytes))
		rawBytes, err  = stub.GetState(acc2);
		if rawBytes == nil {
			return nil, errors.New("Cannot access unknown account " + acc2)
		}
		val2, _ = strconv.Atoi(string(rawBytes))
		val, err = strconv.Atoi(args[2])
		if err != nil {
			return nil, errors.New("Expected integer value, got " + args[2])
		}
		val1 = val1 - val
		val2 = val2 + val
		err = stub.PutState(acc1, []byte(strconv.Itoa(val1)))
		if err != nil {
			return nil, err
		}
		err = stub.PutState(acc2, []byte(strconv.Itoa(val2)))
		return nil, err
	}

	err = errors.New("Unknown function, expected 'open', 'close', 'deposit', 'withdrawal' or 'transfer'")
	return nil, err
}


func (t *ExchangeChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	var acc1 string
	
	if function == "list" {
		if len(args) != 0 {
			return nil, errors.New("Incorrect number of arguments. Expecting 0")
		}
		var iter shim.StateRangeQueryIteratorInterface
		iter, err = stub.RangeQueryState(" ", "~")
		if err != nil {
			return nil, err
		}
		accounts := ""
		for (iter.HasNext()) {
			key, _, _ := iter.Next()
			accounts = accounts + key + "\n"
		}
		err = iter.Close()
		return []byte(accounts), err
	}

	if function == "query" {
		if len(args) != 1 {
			return nil, errors.New("Expected argument: Account")
		}
		acc1 = args[0]
		rawBytes, err := stub.GetState(acc1);
		if rawBytes == nil {
			return nil, errors.New("Unknown account " + acc1)
		}
		return rawBytes, err
	}

	err = errors.New("Unknown function, expected 'query' or 'list'")
	return nil, err
}


func main() {
	err := shim.Start(new(ExchangeChaincode))
	if err != nil {
		fmt.Printf("Error starting ExchangeChaincode: %s", err)
	}
}
