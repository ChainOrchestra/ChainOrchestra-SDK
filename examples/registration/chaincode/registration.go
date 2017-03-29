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
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"chainorchestra.com/perms"
)


type RegistrationChaincode struct {
}


func (t *RegistrationChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
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


func (t *RegistrationChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	guest := args[0]
	
	permissions, err := perms.CheckPermissions(stub)
	if err != nil {
		return nil, err
	}
	
	guestStateBytes, err := stub.GetState(guest)
	if err != nil {
		return nil, err
	}
	
	guestState := string(guestStateBytes)
	if function == "register" {
		if (0 == permissions & perms.DeskClerk) {
			return nil, errors.New("Need Desk Clerk rights to register")
		}
		if guestState != "" {
			return nil, errors.New("Cannot register guest '" + guest + "', status already '" + guestState + "'")
		}
		err = stub.PutState(guest, []byte("registered"))
		return nil, err
	}
	
	if function == "checkin" {
		if (0 == permissions & perms.SecurityGuard) {
			return nil, errors.New("Need Security Guard rights to checkin")
		}
		if guestState != "registered" {
			return nil, errors.New("Expected 'registered' state for guest '" + guest + "', got '" + guestState + "'")
		}
		err = stub.PutState(guest, []byte("checked in"))
		return nil, err
	}
	
	if function == "checkout" {
		if (0 == permissions & perms.SecurityGuard) {
			return nil, errors.New("Need Security Guard rights to checkout")
		}
		if guestState != "checked in" {
			return nil, errors.New("Expected 'checked in' state for guest '" + guest + "', got '" + guestState + "'")
		}
		err = stub.PutState(guest, []byte("registered"))
		return nil, err
	}
	
	if function == "unregister" {
		if (0 == permissions & perms.DeskClerk) {
			return nil, errors.New("Need Desk Clerk rights to unregister")
		}
		if guestState != "registered" {
			return nil, errors.New("Expected 'registered' state for guest '" + guest + "', got '" + guestState + "'")
		}
		err = stub.DelState(guest)
		return nil, err
	}
	
	err = errors.New("Unknown function, expected 'register', 'checkin', 'checkout' or 'unregister'")
	return nil, err
}


func (t *RegistrationChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	
	if function == "list" {
		if len(args) != 0 {
			return nil, errors.New("Incorrect number of arguments. Expecting 0")
		}
		var iter shim.StateRangeQueryIteratorInterface
		iter, err = stub.RangeQueryState(" ", "~")
		if err != nil {
			return nil, err
		}
		guests := ""
		for (iter.HasNext()) {
			key, _, _ := iter.Next()
			guests = guests + key + "\n"
		}
		err = iter.Close()
		return []byte(guests), err
	}

	if function == "query" {
		if len(args) != 1 {
			return nil, errors.New("Expected argument: Account")
		}
		guest := args[0]
		rawBytes, err := stub.GetState(guest)
		if rawBytes == nil {
			return nil, errors.New("Unknown guest '" + guest)
		}
		return rawBytes, err
	}

	err = errors.New("Unknown function, expected 'query' or 'list'")
	return nil, err
}


func main() {
	err := shim.Start(new(RegistrationChaincode))
	if err != nil {
		fmt.Printf("Error starting RegistrationChaincode: %s", err)
	}
}
