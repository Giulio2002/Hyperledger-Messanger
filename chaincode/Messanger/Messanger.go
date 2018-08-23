package main

import (
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
    "fmt"
    "encoding/json"
)

type Messanger struct {
}

// User Definition
type User struct {
    Msgs [] string
}

func (c *Messanger) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}

func (c *Messanger) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

    function, args := stub.GetFunctionAndParameters()

    if function == "sendMessage" {
        if len(args) != 4 {
            return shim.Error("Invalid Arguments Number")
        }
        // init args
        loginArgs := make([][]byte, 3)
        // login object
        loginArgs[0]= []byte("login")
        loginArgs[1]= []byte(args[0])
        loginArgs[2]= []byte(args[1])
        // check response
        loginRes := stub.InvokeChaincode("Authenticator",  loginArgs, "")
        if loginRes.Status != 200 {
            return shim.Error("Error in login")
        }
        // verify if the recipient exist
        existArgs := make([][]byte, 2)
        // login object
        existArgs[0]= []byte("exist")
        existArgs[1]= []byte(args[0])
        // check response
        existRes := stub.InvokeChaincode("Authenticator",  existArgs, "")
        if existRes.Status != 200 {
            return shim.Error("User not found")
        }
        // get user
        v, err := stub.GetState(args[2])
        if err != nil {
            return shim.Error(err.Error())
        }
        // send message
        var user User;
        json.Unmarshal(v, &user)
        user.Msgs = append(user.Msgs, args[0] + ": " + args[3])
        userAsBytes, err := json.Marshal(user)
        if err != nil {
            return shim.Error(err.Error())
        }
        // save it in the network
        err = stub.PutState(args[2], userAsBytes)

        if err != nil {
            return shim.Error(err.Error())
        }
        return shim.Success(nil)


    } else if function == "query" {
        if len(args) != 2 {
            return shim.Error("Invalid Arguments Number")
        }
        // init args
        loginArgs := make([][]byte, 3)
        // login object
        loginArgs[0]= []byte("login")
        loginArgs[1]= []byte(args[0])
        loginArgs[2]= []byte(args[1])
        // check response
        loginRes := stub.InvokeChaincode("Authenticator",  loginArgs, "")
        if loginRes.Status != 200 {
            return shim.Error("Error in login")
        }
        // get user
        v, err := stub.GetState(args[0])
        if err != nil {
            return shim.Error(err.Error())
        }
        // send message
        var user User;
        json.Unmarshal(v, &user)
        i := 0
        retval := ""
        for i < len(user.Msgs) {
            retval = retval + user.Msgs[i] + ", "
            i = i + 1;
        }
        return shim.Success([]byte(retval))
    } else {
        return shim.Error("Invalid Command")
    }

    return shim.Success(nil)
}

func main() {
    err := shim.Start(new(Messanger))
    if err != nil {
        fmt.Printf("Error starting chaincode sample: %s", err)
    }
}
