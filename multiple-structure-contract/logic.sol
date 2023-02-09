// logic.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./data.sol";

contract Logic {
    /**
     * @dev Set a value.
     * @param _data Contract address
     * @param _x value
     */
    function setX(Data _data, uint256 _x) public {
        _data.setX(_x);
    }

    /**
     * @dev Set a value.
     * @param _addr Contract address
     * @param _x value
     */
    function setXFromAddress(address _addr, uint256 _x) public {
        Data data = Data(_addr);
        data.setX(_x);
    }
}
