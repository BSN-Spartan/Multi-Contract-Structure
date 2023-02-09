// receiver.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Receiver {
    uint256 public x;

    /**
     * @dev Set a value.
     * @param _x value
     */
    function setX(uint256 _x) public returns (uint256) {
        x = _x + 1;
        return x;
    }

    /**
     * @dev get x value.
     */
    function getX() public view returns (uint256) {
        return x;
    }
}
