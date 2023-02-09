# Structure of Multi-Contract Calls

[![Smart Contract](https://badgen.net/badge/smart-contract/Solidity/orange)](https://soliditylang.org/) 

It is not a good design for an application to uses only one contract to implement the whole business, because the data carrying will exceed the limits of the block. It is also difficult to upgrade the contract if the data and logic are placed in the same contract. In this case, the structure of multiple contracts should be considered at the beginning of the design. Here we introduce two types of multi-contract calls for beginners.


- Import
- Call and delegatecall

## Prerequisite

Before using a smart contract, it is important to have a basic understanding of Ethereum and Solidity. 

## Overview

### Import

The Logic contract implements the function calls in the Data contract by importing the Data contract. When the `setX` function in the Logic contract is called after inputting the address of the Data contract and the value, the value in the Data contract changes accordingly.

The `setX` and `setXFromAddress` functions implement exactly the same logic but are written in different forms.

This model can be applied to simple business scenarios where data and logic are separated.

### Call and Delegatecall

The `callSetX` function in the Caller contract implements the `call()` function and points to the `setX` function call of a contract. When calling the `callSetX` function, the first input is the address of the Receiver contract and the second input is an integer. Based on the nature of the `call()` function, the `setX` function in the Receiver contract will be called, and the result is returned by adding one to the value of the call state. At this time, the value of variable x in the Recevier contract is **changed**, and the value of x in the Caller contract is **NOT** changed.

The `delegatecallSetX` function in the Caller contract implements the `delegatecall()` function and points to the `setX` function call of a contract. When calling the `delegatecallSetX` function, the first input is the address of the Receiver contract and the second input is an integer. Based on the nature of the `delegatecall()` function, the `setX` function in the Receiver contract will be called, and the result is returned by adding one to the value of the call state. At this time, the value of variable x in the Recevier contract is **NOT** changed, and the value of x in the Caller contract is **changed**.

**Note**: The storage layout of the Caller contract implementing `delegatecall()` function must be the same as the Receiver contract.

The model can be applied to complex scenarios with mixed data and logic.

## Usage

Get the smart contract from [GitHub](https://github.com/BSN-Spartan/NFT-Fractional-Contract/tree/main/contracts), or get the source code by command:

```
$ git clone https://github.com/BSN-Spartan/Multi-Contract-Structure.git
```

For beginners, the contracts in this application can be deployed by the steps in [Spartan Quick Testing](https://www.spartan.bsn.foundation/main/quick-testing#step1).

There are four contracts in this application: data，logic，caller and receiver. Follow steps below to deploy and use them:

1. Deploy data.sol and logic.sol.
2. Get the address of data contract, then call the functions in logic contract and verify the result.
3. Deploy receiver.sol and caller.sol.
4. Get the address of receiver contract, then call the functions in caller contract and verify the result.

## Main Functions

### Import

#### logic.sol: 

```
setX(Data _data, uint _x) public
```

Data will not be returned after calling this function, but you can observe the following results after a successful call:
The value x has been changed in Data contract.

```
setXFromAddress(address _addr, uint _x) public
```

Data will not be returned after calling this function, but you can observe the following results after a successful call:
The value x has been changed in Data contract.

### Call and Ddelegatecall

#### caller.sol: 

```
callSetX(address _addr, uint _x) public
```

Data will not be returned after calling this function, but you can observe the following results after a successful call:

The value x in Recevier contract has been changed while the value x in Caller contract remains the same.

```
delegatecallSetX(address  _addr, uint _x) public
```

Data will not be returned after calling this function, but you can observe the following results after a successful call:

The value x in Recevier contract remains the same while the value x in Caller contract has been changed.

### Query

```
getX() public view returns (uint256)
```

Return the value x by calling this function.

## License

Spartan-NFT Contracts is released under the [Spartan License](https://github.com/BSN-Spartan/NFT/blob/main/LICENSE).
