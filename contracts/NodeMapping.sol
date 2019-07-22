pragma solidity ^0.4.25;
//pragma solidity ^0.5.7;

contract NodeMapping {

    mapping (string => uint) nodeIdMap; 
    mapping (address => uint) nodeAddressMap; 
    mapping (string => uint) nodeIpMap; 
    address owner;
    address operator1;
    address operator2;
    address operator3;
    address operator4;
    uint256 nodeCount;

    constructor() public {
        owner = msg.sender;
        operator1 = msg.sender;
        operator2 = msg.sender;
        operator3 = msg.sender;
        operator4 = msg.sender;
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
    function SetOperators(address newOperator1, address newOperator2, address newOperator3, address newOperator4) OwnerOrOperator public {
        operator1 = newOperator1;
        operator2 = newOperator2;
        operator3 = newOperator3;
        operator4 = newOperator4;
    }
    function GetOperator1() public view returns (address) {
        return operator1;
    }
    function GetOperator2() public view returns (address) {
        return operator2;
    }
    function GetOperator3() public view returns (address) {
        return operator3;
    }
    function GetOperator4() public view returns (address) {
        return operator4;
    }
    function GetOwner() public view returns (address) {
        return owner;
    }
    function GetNodeCount() public view returns (uint256) {
        return nodeCount;
    }
    modifier OwnerOrOperator {
        require(msg.sender == owner || msg.sender == operator1 || msg.sender == operator2 || msg.sender == operator3 || msg.sender == operator4);
        _;
    }
}
