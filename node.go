package main

import (
    "flag"
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"
    "bufio"
    "os"
    "strings"
    "syscall"
    "os/signal"
    "encoding/hex"
    "time"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/glendc/go-external-ip"
)

const logo = `
|
|____  __
|\   \/  /___________  ____   _____
| \     // __ \_  __ \/  _ \ /     \
| /     \  ___/|  | \(  <_> )  Y Y  \
|/___/\  \___  >__|   \____/|__|_|  /
|      \_/   \/                   \/
|             Node Deployment System
`

type NodeType struct {
    Name                  string
    RequiredCollateral    int
    ContractAddress       string
}

var NodeTypes = map[int]NodeType {
    1 : NodeType{
                Name:                  "Chain Node",
                RequiredCollateral:    5000,
                ContractAddress:       "0x3717AD55666577Eb92fCa3e5F9F71958bD60c620",
                //ContractAddress:       "0xC0968a743cF57a61baC462FF903d9318B61426f6",
    },
    2 : NodeType{
                Name:                  "Xero Node",
                RequiredCollateral:    20000,
                ContractAddress:       "0xc46Cc53b8F09fe6F4eB6b6dF8AD5c6Fe5DA6638B",
    },
    3 : NodeType{
                Name:                  "Link Node",
                RequiredCollateral:    40000,
                ContractAddress:       "0xE44389C26FDEb581dEa7Df91Efd0665a7cd404c1",
    },
    4 : NodeType{
                Name:                  "Super Node",
                RequiredCollateral:    80000,
                ContractAddress:       "0x93B7a5c74793DCba765a1dD163e1744622306651",
    },
}

func main() {
    // Flags
    adminFlag := flag.Bool("admin", false, "a bool")
    flag.Parse()

    // Print logo
    fmt.Println(logo)
    // Setup graceful exit
    var gracefulStop = make(chan os.Signal)
    signal.Notify(gracefulStop, syscall.SIGTERM)
    signal.Notify(gracefulStop, syscall.SIGINT)
    go func() {
           <-gracefulStop
           fmt.Printf("\nExiting Program\n")
           os.Exit(0)
    }()

    var selectionFlag = false

    for selectionFlag != true {

        var contractOption int

        fmt.Println("1) Add a New Node")
        fmt.Println("2) Remove an Existing Node")
        fmt.Println("3) Lookup Existing Node")
        fmt.Println("4) Exit")
        if *adminFlag {
            fmt.Println("5) Deploy Contract")
        }

        _, _ = fmt.Scan(&contractOption)

        if contractOption == 1 {

            nodeType := getNodeType()

            if nodeType != 5 {

                if checkBinExistence() {

                    reader := bufio.NewReader(os.Stdin)

                    // Get Node ID
                    nodeId := hex.EncodeToString(getNodeId())
                    fmt.Println("\nNode ID Found: " + nodeId)

                    // Get Node IP
                    var nodeIp string
                    consensus := externalip.DefaultConsensus(nil, nil)
                    ip, _ := consensus.ExternalIP()
                    nodeIp = ip.String()

                    fmt.Println("Node IP Address Found: " + nodeIp)

                    // Get Node Port
                    var nodePort string
                    fmt.Println("Enter Node Port:")
                    nodePort, _ = reader.ReadString('\n')
                    nodePort = strings.TrimSuffix(nodePort, "\n")

                    // Get Private Key
                    var privateKey string
                    fmt.Println("Enter Private Key of Address Containing Node Collateral:")
                    privateKey, _ = reader.ReadString('\n')
                    privateKey = strings.TrimSuffix(privateKey, "\n")

                    addNode(nodeId, nodeIp, nodePort, privateKey, nodeType)

                    selectionFlag = true

                } else {
                    fmt.Println("Node Not Found - Unable To Continue")
                    os.Exit(0)
                }

            } else {

                selectionFlag = false

            }

        } else if contractOption == 2 {

            nodeType := getNodeType()

            if nodeType != 5 {

                reader := bufio.NewReader(os.Stdin)

                // Get Private Key
                var privateKey string
                fmt.Println("Enter Private Key To Release Collateral and Delete Node:")
                privateKey, _ = reader.ReadString('\n')
                privateKey = strings.TrimSuffix(privateKey, "\n")

                removeNode(privateKey, nodeType)

                selectionFlag = true

           } else {

               selectionFlag = false

           }

        } else if contractOption == 3 {

            nodeType := getNodeType()

            if nodeType != 5 {

                reader := bufio.NewReader(os.Stdin)

                // Get Private Key
                var nodeAddress string
                fmt.Println("Enter Node Address To Lookup Node:")
                nodeAddress, _ = reader.ReadString('\n')
                nodeAddress = strings.TrimSuffix(nodeAddress, "\n")

                checkNodeExistence(nodeAddress, nodeType)

                selectionFlag = true

           } else {

               selectionFlag = false

           }

        } else if contractOption == 4 {

            selectionFlag = true
            fmt.Printf("\nExiting Program\n")
            os.Exit(0)

        } else if contractOption == 5 && *adminFlag {

            reader := bufio.NewReader(os.Stdin)

            // Get Private Key
            var deploymentKey string
            fmt.Println("Enter Private Key For Contract Deployment:")
            deploymentKey, _ = reader.ReadString('\n')
            deploymentKey = strings.TrimSuffix(deploymentKey, "\n")

            // Deploy Contracts
            contractDeployment(deploymentKey)
            selectionFlag = true

        } else {

            fmt.Println("\nInvalid Input\n")
            selectionFlag = false

        }

    }
}

func getNodeType() int {

    var selectionFlag = false
    var nodeSelection int

    for selectionFlag != true {

        fmt.Println("\n")
        fmt.Println("1) Chain Node")
        fmt.Println("2) Zero Node")
        fmt.Println("3) Link Node")
        fmt.Println("4) Super Node")
        fmt.Println("5) Main Menu")

        _, _ = fmt.Scan(&nodeSelection)

        if nodeSelection == 1 {

            selectionFlag = true
            return nodeSelection

        } else if nodeSelection == 2 {

            selectionFlag = true
            return nodeSelection

        } else if nodeSelection == 3 {

            selectionFlag = true
            return nodeSelection

        } else if nodeSelection == 4 {

            selectionFlag = true
            return nodeSelection

        } else if nodeSelection == 5 {

            selectionFlag = true
            return nodeSelection

        } else {

            fmt.Println("\nInvalid Input\n")
            selectionFlag = false

        }

    }
    return nodeSelection
}

func addNode(nodeId string, nodeIp string, nodePort string, key string, nodeType int) {
    //homeDirectory := getHomeDirectory()
    //homeDirectory := getHomeDirectory()

    client, err := ethclient.Dial(DefaultDataDir() + "/geth.ipc")
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA(key)
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = new(big.Int).Mul(big.NewInt(int64(NodeTypes[nodeType].RequiredCollateral)), big.NewInt(1e+18)) // in wei
    auth.GasLimit = uint64(3000000) // in units
    auth.GasPrice = gasPrice

    address := common.HexToAddress(NodeTypes[nodeType].ContractAddress)
    instance, err := NewNodeContract(address, client)
    if err != nil {
        log.Fatal(err)
    }

    // Add node
    tx, err := instance.AddNode(auth, nodeId, nodeIp, nodePort)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Adding " + NodeTypes[nodeType].Name)
    fmt.Printf("Tx Sent: %s", tx.Hash().Hex())
    fmt.Println("\n")

}

func removeNode(key string, nodeType int) {
    //homeDirectory := getHomeDirectory()

    client, err := ethclient.Dial(DefaultDataDir() + "/geth.ipc")
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA(key)
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0) // in wei
    auth.GasLimit = uint64(3000000) // in units
    auth.GasPrice = gasPrice

    address := common.HexToAddress(NodeTypes[nodeType].ContractAddress)
    instance, err := NewNodeContract(address, client)
    if err != nil {
        log.Fatal(err)
    }

    // Remove node
    tx, err := instance.RemoveNode(auth)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Removing " + NodeTypes[nodeType].Name)
    fmt.Printf("Tx sent: %s", tx.Hash().Hex())
    fmt.Println("\n")

}

func checkNodeExistence(nodeAddress string, nodeType int) {
    //homeDirectory := getHomeDirectory()

    client, err := ethclient.Dial(DefaultDataDir() + "/geth.ipc")
    if err != nil {
        log.Fatal(err)
    }

    address := common.HexToAddress(NodeTypes[nodeType].ContractAddress)
    instance, err := NewNodeContract(address, client)
    if err != nil {
        log.Fatal(err)
    }

    // Get node id
    id, err := instance.GetNodeId(&bind.CallOpts{}, common.HexToAddress(nodeAddress))
    if err != nil {
        log.Fatal(err)
    }

    // Get node ip
    ip, err := instance.GetNodeIp(&bind.CallOpts{}, common.HexToAddress(nodeAddress))
    if err != nil {
        log.Fatal(err)
    }

    // Get node port
    port, err := instance.GetNodePort(&bind.CallOpts{}, common.HexToAddress(nodeAddress))
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Checking Existence of " + NodeTypes[nodeType].Name)
    if id != "" && ip != "" && port != "" {
        fmt.Println("Node Id: " + id)
        fmt.Println("Node Ip: " + ip)
        fmt.Println("Node Port: " + port)
    } else {
        fmt.Println("Node Not Found")
    }
    fmt.Println("\n")
}
/*
func checkBinExistence() bool {
    var nodeExistenceFlag = false
    var nodeInstallationFlag = false

    for nodeExistenceFlag == false {
        nodeExistenceFlag = getBlockHeight()
        if nodeExistenceFlag {
            return true
        } else if nodeInstallationFlag == false {
            reader := bufio.NewReader(os.Stdin)
            var installNodeString string
            fmt.Println("Would You Like To Install Node Binary? (Y/N)")
            installNodeString, _ = reader.ReadString('\n')
            installNodeString = strings.TrimSuffix(installNodeString, "\n")
                if installNodeString == "Y" {
                    nodeInstallationFlag = installNode()
                } else {
                    fmt.Println("\nUnable To Find Active Node - Exiting Dashboard")
                    os.Exit(0)
                }
        }
        time.Sleep(20 * time.Second)
    }
    return false
}

func installNode() bool {
    exec.Command("/bin/sh", "/home/nucleos/scripts/node.sh")
    return true
}

// Get lastes block and check if node is active
func getBlockHeight() bool {
    homeDirectory := getHomeDirectory()

    client, err := ethclient.Dial(homeDirectory + "")
    if err != nil {
        fmt.Println("Xero Node Not Found")
        return false
    }

    _, err = client.BlockByNumber(context.Background(), nil)
    if err != nil {
        fmt.Println("Xero Node Not Found")
        return false
    }
    fmt.Println("Xero Node Found")
    return true
}
*/
func contractDeployment(key string) {
    for _, nodeType := range NodeTypes {
        time.Sleep(30 * time.Second)
        deployNodeTypeContract(key, nodeType.RequiredCollateral)
    }

    time.Sleep(30 * time.Second)
    mappingAddress := deployMappingContract(key)

    count := int(0)
    for _, _ = range NodeTypes {
        time.Sleep(30 * time.Second)
        updateNodeTypeContractMappingAddress(key, count, mappingAddress)
        count++
    }

}

func updateNodeTypeContractMappingAddress(key string, nodeType int, mappingAddress common.Address) {
    //homeDirectory := getHomeDirectory()
    //homeDirectory := getHomeDirectory()

    client, err := ethclient.Dial(DefaultDataDir() + "/geth.ipc")
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA(key)
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0) // in wei
    auth.GasLimit = uint64(3000000) // in units
    auth.GasPrice = gasPrice

    address := common.HexToAddress(NodeTypes[nodeType].ContractAddress)
    instance, err := NewNodeContract(address, client)
    if err != nil {
        log.Fatal(err)
    }

    // Add node
    tx, err := instance.UpdateNodeMappingAddress(auth, mappingAddress)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Updating Mapping Address: " + NodeTypes[nodeType].Name)
    fmt.Printf("Tx Sent: %s", tx.Hash().Hex())
    fmt.Println("\n")

}

func deployNodeTypeContract(key string, contractCollateral int) {
    //homeDirectory := getHomeDirectory()

    client, err := ethclient.Dial(DefaultDataDir() + "/geth.ipc")
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA(key)
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0) // in wei
    auth.GasLimit = uint64(3000000) // in units
    auth.GasPrice = gasPrice

    // Deploy contract
    address, tx, _, err := DeployNodeContract(auth, client, big.NewInt(int64(contractCollateral)))
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Deploying Contract...\n\n")
    fmt.Printf("Tx sent: %s\n", tx.Hash().Hex())
    fmt.Printf("Contract Address: %s", address.Hex())
    fmt.Println("\n")

    f, err := os.OpenFile("contractDeployments.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    if _, err := f.Write([]byte(address.Hex() + "\n")); err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
         log.Fatal(err)
    }
    fmt.Println("Saving Contract Deployment Output...\n")

}

func deployMappingContract(key string) common.Address {
    //homeDirectory := getHomeDirectory()

    client, err := ethclient.Dial(DefaultDataDir() + "/geth.ipc")
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA(key)
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0) // in wei
    auth.GasLimit = uint64(3000000) // in units
    auth.GasPrice = gasPrice

    // Deploy contract
    address, tx, _, err := DeployNodeMappingContract(auth, client)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Deploying Contract...\n\n")
    fmt.Printf("Tx sent: %s\n", tx.Hash().Hex())
    fmt.Printf("Contract Address: %s", address.Hex())
    fmt.Println("\n")

    f, err := os.OpenFile("contractDeployments.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    if _, err := f.Write([]byte(address.Hex() + "\n")); err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
         log.Fatal(err)
    }
    fmt.Println("Saving Contract Deployment Output...\n")

    return address
}
/*
// Get user home directory from env
func getHomeDirectory() string {
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    return usr.HomeDir
}

// Retrieve nodekey and calculate enodeid
func getNodeId() []byte {
    b, err := ioutil.ReadFile(getHomeDirectory() + "/.xerom/geth/nodekey")
    if err != nil {
        fmt.Print(err)
        return []byte{}
    }
    enodeId, err := crypto.HexToECDSA(string(b))
    if err != nil {
        fmt.Print(err)
        return []byte{}
    }
    pubkeyBytes := crypto.FromECDSAPub(&enodeId.PublicKey)[1:]
    return pubkeyBytes
}
*/
