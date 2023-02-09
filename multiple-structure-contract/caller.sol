// caller.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Caller {
    uint256 public x;

    // event for EVM logging
    event Response(bool success, bytes data);

    /**
     * @dev Call the setX method in the contract for the given contract address.
     * @param _addr contract address
     * @param _x value
     */
    function callSetX(address _addr, uint256 _x) public {
        (bool success, bytes memory data) = _addr.call(
            abi.encodeWithSignature("setX(uint256)", _x)
        );

        emit Response(success, data);
    }

    /**
     * @dev Delegatecall the setX method in the contract for the given contract address.
     * @param _addr contract address
     * @param _x value
     */
    function delegatecallSetX(address _addr, uint256 _x) public {
        (bool success, bytes memory data) = _addr.delegatecall(
            abi.encodeWithSignature("setX(uint256)", _x)
        );

        emit Response(success, data);
    }

    /**
     * @dev get x value.
     */
    function getX() public view returns (uint256) {
        return x;
    }
}
