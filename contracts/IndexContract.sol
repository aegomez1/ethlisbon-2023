// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

contract IndexContract {
    struct SubmissionData {
        address contractAddress;
        string ipfsCID;
    }

    mapping(address => SubmissionData) public submissionMap;

    // Declare an event
    event SubmissionAdded(address indexed contractAddress, string ipfsCID);

    function addSubmission(address contractAddr, string memory cid) public {
        submissionMap[contractAddr] = SubmissionData(contractAddr, cid);


        // Emit the SubmissionAdded event
        emit SubmissionAdded(contractAddr, cid );
    }
}
