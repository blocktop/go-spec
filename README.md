# go-spec
> The interface specifications for the blocktop blockchain development kit components.

## Overview

This repository contains the interfaces that components must implement in order to participate in a blocktop-based project. One can see the relationships between these components by examining tbe interfaces.

## Interfaces

### Controller

A Controller manages blockchains and the flow of Vlocks and Transactions into them.

### Blockchain

A Blockchain employs a BlockGenerator to creaye new blocks locally. It runs a Consensus component which tells it which head Block on which the next Block should generated. It also receives newly committed Blocks from Consensus and passes them to the BlockGenerator for computation and storage.

### Consensus

A Consensus component tracks incoming Blocks, whether generated locally or received from another NetworkNode. When Blocks that have the same parent Block arrive, it uses a BlockComparator function to determine the Block that is to remain in competition. Consensus confirms Blocks and passes them to Blockchain for further processing.

### BlockGenerator

A BlockGenerator is responsible for the generation of new Blocks, including attaching transactions to them. It may use any algorithm required for the syatem being implemented, including hash generation (Proof-of-Work), coin staking (Proof-of-Stake), or others. A BlockGenerator also maintains a collection of TransactionHandlers. A BlockGenerator is responsible for validating Blocks and committing them to storage.

### TransactionHandler

A TransaxtionHandler validates and executes Transaxtions.

### Block

A Block is the unit member of a Blockchain. It has a pointer to its parent and holds a collection of Transactions. A Block conforms to the Marshalled interface which allows it to be serialized and deserialized for network communication.

### Transaction

A Transaction is an instruction and holds a set of parameters tha

a TransactionHandler uses to co.pute the intended update to the state of the system.
e
