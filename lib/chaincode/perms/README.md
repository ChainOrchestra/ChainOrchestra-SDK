# [ChainOrchestra-SDK](https://github.com/ChainOrchestra/ChainOrchestra-SDK) 

## Chaincode permissions module

### Permissions module general design

#### Intended use

The rationale for this library is to have a very lightweight module, consistent accross all chaincode applications,
that doesn't re-implement the cryptographic services already in place with Hyperledger. It just checks rights and
permissions granted to already enrolled users.

As of Hyperledger v0.6, it uses the position member services attribute with a bitmask to set prioritized permissions.

The perms.go library must be included by all chaincode applications that need to implement some level of resource sharing.
The module defines a set of roles that are granted to blockchain users, according to which the chaincode will let the user
perform certain actions.

#### Example implementation

A banking application defines 3 positions: Bank Clerk, Cashier and Accountant. 
The Accountant is granted both BankClerk and Cashier rights.

```
const BankClerk			uint64 = 0x0000000000100000
const Cashier			uint64 = 0x0000000000200000
const Accountant		uint64 = 0x0000000000300000
```

System administrators and power users have cross-functional rights that include those defined above.

```
const SystemAdmin		uint64 = 0xffffffffffff000f
const PowerUser			uint64 = 0x0000000000ff0000
```


### [perms.go](perms.go) library usage


  * Importing the library from user chaincode.

```
import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"chainorchestra.com/perms"
)
```

  * Using the ```shim.ChaincodeStubInterface``` to get permissions mask.

```
permissions, err := perms.CheckPermissions(stub)
```

  * Checking the permissions mask for a specific grant.

```
if (0 == permissions & perms.BankClerk) {
	return nil, errors.New("Need Bank Clerk rights to open or close account")
}
```

