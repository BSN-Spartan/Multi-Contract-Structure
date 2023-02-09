// data.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Data {
    uint256 public x;

    /**
     * @dev Set a value.
     * @param _x value
     */
    function setX(uint256 _x) public returns (uint256) {
        x = _x;
        return x;
    }

    /**
     * @dev get x value.
     */
    function getX() public view returns (uint256) {
        return x;
    }
}
