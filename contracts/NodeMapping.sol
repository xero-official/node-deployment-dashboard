pragma solidity ^0.4.25;
//pragma solidity ^0.5.7;

contract NodeMapping {

    mapping (string => string) nodeIdMap; 
    mapping (address => address) nodeAddressMap; 
    mapping (string => string) nodeIpMap; 
    address owner;
    address operator;

    constructor(uint setCollateralRequirement) public {
        owner = msg.sender;
        operator = msg.sender;
    }
    function AddNode(address nodeAddress, string memory id, string memory ip) OwnerOrOperator public {
        nodeIdMap[id] = id;
        nodeAddressMap[nodeAddress] = nodeAddress;
        nodeIpMap[ip] = ip;
    }
    function RemoveNode(address nodeAddress, string memory id, string memory ip) OwnerOrOperator public {
        delete nodeIdMap[id];
        delete nodeAddressMap[nodeAddress];
        delete nodeIpMap[ip];
    }
    function CheckExistence(address nodeAddress, string memory id, string memory ip) public returns (bool) {
        assert(keccak256(abi.encodePacked(nodeIdMap[id])) == keccak256(abi.encodePacked(id)) || nodeAddressMap[nodeAddress] == nodeAddress || keccak256(abi.encodePacked(nodeIpMap[ip])) == keccak256(abi.encodePacked(ip)));
        return true;    
    }
    function SetOperator(address newOperator) OwnerOrOperator public {
        operator = newOperator;
    }

    modifier OwnerOrOperator {
        require(msg.sender == owner || msg.sender == operator);
        _;
    }
}
