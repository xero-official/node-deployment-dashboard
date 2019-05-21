package main

import (
    "os/user"
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

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
}

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

// Temp global vars for testing
var (
    collateral = 5000
    contractAddress = "0xe44389c26fdeb581dea7df91efd0665a7cd404c1"
)

func main() {
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
        fmt.Println("3) Exit")

        _, _ = fmt.Scan(&contractOption)

        if contractOption == 1 {

            // Need to add node type selection here

            reader := bufio.NewReader(os.Stdin)

            // Get Node ID
            var nodeId string
            fmt.Println("Enter Node ID:")
            nodeId, _ = reader.ReadString('\n')
            nodeId = strings.TrimSuffix(nodeId, "\n")

            // Get Node IP
            var nodeIp string
            fmt.Println("Enter Node IP Address:")
            nodeIp, _ = reader.ReadString('\n')
            nodeIp = strings.TrimSuffix(nodeIp, "\n")

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

            addNode(nodeId, nodeIp, nodePort, privateKey)

            selectionFlag = true

        } else if contractOption == 2 {

            reader := bufio.NewReader(os.Stdin)

            // Get Private Key
            var privateKey string
            fmt.Println("Enter Private Key To Release Collateral and Delete Node:")
            privateKey, _ = reader.ReadString('\n')
            privateKey = strings.TrimSuffix(privateKey, "\n")

            removeNode(privateKey)

            selectionFlag = true

        } else if contractOption == 3 {

            selectionFlag = true
            fmt.Printf("\nExiting Program\n")
            os.Exit(0)

        } else {

            fmt.Println("\nInvalid Input\n")
            selectionFlag = false

        }

    }
}

func addNode(nodeId string, nodeIp string, nodePort string, key string) {
    homeDirectory := getHomeDirectory()

    client, err := ethclient.Dial(homeDirectory + "/.xerom/geth.ipc")
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
    auth.Value = new(big.Int).Mul(big.NewInt(int64(collateral)), big.NewInt(1e+18)) // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

    address := common.HexToAddress(contractAddress)
    instance, err := NewNodeContract(address, client)
    if err != nil {
        log.Fatal(err)
    }

    // Add node
    tx, err := instance.AddNode(auth, nodeId, nodeIp, nodePort)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s", tx.Hash().Hex())
    fmt.Println("\n")

}

func removeNode(key string) {
    homeDirectory := getHomeDirectory()

    client, err := ethclient.Dial(homeDirectory + "/.xerom/geth.ipc")
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
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

    address := common.HexToAddress(contractAddress)
    instance, err := NewNodeContract(address, client)
    if err != nil {
        log.Fatal(err)
    }

    // Remove node
    tx, err := instance.RemoveNode(auth)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s", tx.Hash().Hex())
    fmt.Println("\n")

}

func getHomeDirectory() string {
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    return usr.HomeDir
}
