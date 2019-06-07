pragma solidity ^0.4.25;
//pragma solidity ^0.5.7;

contract NodeContract {
    struct Node {
        uint index;
        address nodeAddress;
        string nodeId;
        string nodeIp;
        string nodePort;
        uint collateral;
    }

    mapping (address => Node) nodeMap; 
    mapping (uint => address) nodeIndexMap; 
    uint nodeCount;
    address owner;
    uint internal requiredCollateral;

    constructor(uint setCollateralRequirement) public {
        nodeCount = 0;
        requiredCollateral = setCollateralRequirement;
        owner = msg.sender;
    }
    function AddNode(string memory id, string memory ip, string memory port) public payable returns (bool) {
        assert(msg.value == requiredCollateral * 1 ether);
        Node storage newNode = nodeMap[msg.sender];
        newNode.nodeAddress = msg.sender;
        newNode.collateral = msg.value;
        newNode.nodeId = id;
        newNode.nodeIp = ip;
        newNode.nodePort = port;
        newNode.index = nodeCount;
        nodeIndexMap[nodeCount] = msg.sender;
        nodeCount++;
        return true;
    }
    function RemoveNode() public returns (bool) {
        assert(nodeMap[msg.sender].collateral == requiredCollateral * 1 ether);
        msg.sender.transfer(nodeMap[msg.sender].collateral);
        nodeMap[nodeIndexMap[nodeCount-1]].index = nodeMap[msg.sender].index;
        nodeIndexMap[nodeMap[msg.sender].index] = nodeIndexMap[nodeCount-1];
        delete nodeIndexMap[nodeCount-1];
        delete nodeMap[msg.sender];
        nodeCount--;
        return true;
    }
    function GetNodeId(address nodeAddress) public view returns (string) {
        return nodeMap[nodeAddress].nodeId;
    }
    function GetNodeIp(address nodeAddress) public view returns (string) {
        return nodeMap[nodeAddress].nodeIp;
    }
    function GetNodePort(address nodeAddress) public view returns (string) {
        return nodeMap[nodeAddress].nodePort;
    }
}