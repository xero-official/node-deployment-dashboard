pragma solidity ^0.4.25;
//pragma solidity ^0.5.7;

contract NodeMapping {

    mapping (string => uint) nodeIdMap; 
    mapping (address => uint) nodeAddressMap; 
    mapping (string => uint) nodeIpMap; 
    address owner;
    address operator;
    uint256 nodeCount;

    constructor(uint setCollateralRequirement) public {
        owner = msg.sender;
        operator = msg.sender;
    }
    function AddNode(address nodeAddress, string memory id, string memory ip) OwnerOrOperator public {
        nodeIdMap[id] = 1111;
        nodeAddressMap[nodeAddress] = 1111;
        nodeIpMap[ip] = 1111;
        nodeCount++;
    }
    function RemoveNode(address nodeAddress, string memory id, string memory ip) OwnerOrOperator public {
        delete nodeIdMap[id];
        delete nodeAddressMap[nodeAddress];
        delete nodeIpMap[ip];
        nodeCount--;
    }
    function CheckExistence(address nodeAddress, string memory id, string memory ip) public returns (bool) {
        assert(nodeIdMap[id] == 1111 || nodeAddressMap[nodeAddress] == 1111 || nodeIpMap[ip] == 1111);
        return true;    
    }
    function SetOperator(address newOperator) OwnerOrOperator public {
        operator = newOperator;
    }
    function GetOperator() public returns (address) {
        return operator;
    }
    function GetOwner() public returns (address) {
        return owner;
    }
    function GetNodeCount() public returns (uint256) {
        return nodeCount;
    }
    modifier OwnerOrOperator {
        require(msg.sender == owner || msg.sender == operator);
        _;
    }
}
