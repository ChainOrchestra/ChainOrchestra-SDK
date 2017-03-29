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

package perms

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const SoftwareEngineer	uint64 = 0xffffffffffffffff
const SystemOperator	uint64 = 0x0000000000000001
const SystemAdmin		uint64 = 0xffffffffffff000f
const PowerUser			uint64 = 0x0000000000ff0000
const DeskClerk			uint64 = 0x0000000000010000
const SecurityGuard		uint64 = 0x0000000000020000
const BankClerk			uint64 = 0x0000000000100000
const Cashier			uint64 = 0x0000000000200000
const Accountant		uint64 = 0x0000000000300000

func CheckPermissions(stub shim.ChaincodeStubInterface) (uint64, error) {
	var err error
	var isOk bool
	var permissions uint64 = 0;

	_, err = stub.ReadCertAttribute("position")
	if err != nil {
		return permissions, err
	}

	isOk, _ = stub.VerifyAttribute("position", []byte("Software Engineer"))
	if (isOk) {
		permissions |= SoftwareEngineer
	}
	isOk, _ = stub.VerifyAttribute("position", []byte("System Operator"))
	if (isOk) {
		permissions |= SystemOperator
	}
	isOk, _ = stub.VerifyAttribute("position", []byte("System Admin"))
	if (isOk) {
		permissions |= SystemAdmin
	}
	isOk, _ = stub.VerifyAttribute("position", []byte("Power User"))
	if (isOk) {
		permissions |= PowerUser
	}
	isOk, _ = stub.VerifyAttribute("position", []byte("Desk Clerk"))
	if (isOk) {
		permissions |= DeskClerk
	}
	isOk, _ = stub.VerifyAttribute("position", []byte("Security Guard"))
	if (isOk) {
		permissions |= SecurityGuard
	}
	isOk, _ = stub.VerifyAttribute("position", []byte("Bank Clerk"))
	if (isOk) {
		permissions |= BankClerk
	}
	isOk, _ = stub.VerifyAttribute("position", []byte("Cashier"))
	if (isOk) {
		permissions |= Cashier
	}
	isOk, _ = stub.VerifyAttribute("position", []byte("Accountant"))
	if (isOk) {
		permissions |= Accountant
	}
	return permissions, err;
}
