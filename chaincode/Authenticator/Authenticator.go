package main

import (
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
    "encoding/json"
)

type Authenticator struct {}

// User Definition
type User struct {
    Name string
    Password string
}

// Implement Init
func(c * Authenticator) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}

// Implement Invoke
func(c * Authenticator) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

    function,
    args: = stub.GetFunctionAndParameters()

    switch function {

        case "login":
            return c.login(stub, args)

        case "register":
            return c.register(stub, args)
        case "exist":
            return c.exist(stub, args)
    }
    return shim.Error("Invalid Command")
}

func(c * Authenticator) register(stub shim.ChaincodeStubInterface, args[] string) pb.Response {

    if len(args) != 2 {
        return shim.Error("Invalid Argument Number")
    }
    // check the password length
    if len(args[1]) < 6 {
        return shim.Error("Password too short")
    }
    // check if theuser is in the network
    v, err: = stub.GetState(args[0])
    var tmp User
    json.Unmarshal(v, & tmp)
    if len(tmp.Password) >= 6 {
        return shim.Error("User already created")
    }
    // create user
    user: = User {
        args[0], args[1]
    }
    userAsBytes, err: = json.Marshal(user)
    if err != nil {
        return shim.Error(err.Error())
    }
    // deploy it in the network
    err = stub.PutState(user.Name, userAsBytes)

    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
}

func(c * Authenticator) login(stub shim.ChaincodeStubInterface, args[] string) pb.Response {
    // check arguments
    if len(args) != 2 {
        return shim.Error("Invalid Argument Number")
    }
    // get user
    value, err: = stub.GetState(args[0])
    if err != nil {
        return shim.Error("User not found")
    }
    // decrypt user[]byte("Status: " + asset.Status)
    var user User
    json.Unmarshal(value, & user)
    // check if the password is correct
    if user.Password != args[1] {
        return shim.Error("Credentials not correct")
    }
    // get asset info
    return shim.Success(nil)
}

func(c * Authenticator) exist(stub shim.ChaincodeStubInterface, args[] string) pb.Response {
    // check arguments
    if len(args) != 1 {
        return shim.Error("Invalid Argument Number")
    }
    // get user
    value, _: = stub.GetState(args[0])
    // decrypt user
    var user User
    json.Unmarshal(value, & user)
    // check if the user exist
    if len(user.Password) < 6 {
        return shim.Error("User not found")
    }
    // get asset info
    return shim.Success(nil)
}

func main() {
    err: = shim.Start(new(Authenticator))
    if err != nil {
        fmt.Printf("Error starting chaincode sample: %s", err)
    }
}
