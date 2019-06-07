package main

import (
    "time"
    "io/ioutil"
    "os/user"
    "os/exec"
    "context"
    "fmt"
    "log"
    "bufio"
    "os"
    "strings"

    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

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
                    os.Exit(0)
                } else {
                    fmt.Println("\nUnable To Find Active Node - Exiting Dashboard")
                    os.Exit(0)
                }
        }
        time.Sleep(20 * time.Second)
    }
    return false
}

// Initiate node installation process
func installNode() bool {
    installSystemService()
    //exec.Command("/bin/sh", "")
    return true
}

// Install systemd service for node
func installSystemService() {
    //cmd := exec.Command("sudo ls > /dev/null")
    //reader := bufio.NewReader(os.Stdin)

    path := "/etc/systemd/system/xeronode.service"
    //file, _ := os.Create(path)
    usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }
    var systemService = "[Unit]\nDescription=XERO Node\nAfter=network.target\n[Service]\nUser=" + usr.Username + "\nGroup=" + usr.Username + "\nType=simple\nRestart=always\nExecStart=/usr/sbin/geth --syncmode=fast --cache=512 --node --datadir=" + usr.HomeDir + "/.xerom\n[Install]\nWantedBy=default.target"
    //fmt.Fprintf(file, "%v", systemService)
    fmt.Println("\nCreating " + path)
    time.Sleep(1 * time.Second)
    //cmd := exec.Command("/bin/sh", "sudo echo " + systemService + " > " + path)
    exec.Command("/bin/sh", "echo " + systemService + " > xeronode.service")
    fmt.Println("Enabling " + path)
    time.Sleep(1 * time.Second)

    fmt.Println("Starting " + path)
    time.Sleep(1 * time.Second)

}

// Get lastes block and check if node is active
func getBlockHeight() bool {
    homeDirectory := getHomeDirectory()

    client, err := ethclient.Dial(homeDirectory + "/.xerom/geth.ipc")
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