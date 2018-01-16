package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("example_chaincode")

type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response  {
	logger.Info("########### example_chaincode Init ###########")
	_, args := stub.GetFunctionAndParameters()
	var A, B string    // Entities
	A = args[0]
	B = args[1]
	err := stub.PutState(A, []byte(B))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}



func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_chaincode Init ###########")
	function, args := stub.GetFunctionAndParameters()

	var A, B string    // Entities
	A = args[0]
	

	if function == "delete" {
		err := stub.DelState(A)
		if err != nil {
			return shim.Error("Failed to delete state")
		}

		return shim.Success(nil)
	}

	if function == "query" {
		Avalbytes, err := stub.GetState(A)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
			return shim.Error(jsonResp)
		}
		return shim.Success(Avalbytes)
	}
	if function == "setdata" {
		B = args[1]
		err := stub.PutState(A, []byte(B))
		if err != nil {
			return shim.Error(err.Error())
		}	
		return shim.Success(nil)
	}

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	return shim.Error("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v"+ args[0])
}


func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
