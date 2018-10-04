# go-spec
> The interface specifications for the blocktop blockchain development kit components.

## Overview

This repository contains the interfaces that components must implement in order to participate in a blocktop-based project. One can see the relationships between these components by examining tbe interfaces.

## Interfaces

### Controller

A `Controller` manages `Blockchain`s and the flow of `Block`s and `Transaction`s into them.

### Blockchain

A `Blockchain` employs a `BlockGenerator` to create new `Block`s locally. It runs a `Consensus` component that tells it which head `Block` on which the next `Block` should be generated. It also receives newly committed `Block`s from `Consensus` and passes them to the `BlockGenerator` for computation and storage.

### Consensus

A `Consensus` component tracks incoming `Block`s, whether generated locally or received from another `NetworkNode`. When `Block`s that have the same parent `Block` arrive, It uses a `BlockComparator` function to determine the `Block` that is to remain in competition. `Consensus` confirms `Block`s and passes them to `Blockchain` for further processing.

### BlockGenerator

A `BlockGenerator` is responsible for the generation of new `Block`s, including attaching transactions to them. It may use any algorithm required for the syatem being implemented, including hash generation (Proof-of-Work), coin staking (Proof-of-Stake), or others. A `BlockGenerator` also maintains a collection of `TransactionHandler`s. A `BlockGenerator` is responsible for validating `Block`s and committing them to storage.

### TransactionHandler

A `TransactionHandler` validates and executes `Transaction`s.

### Block

A `Block` is the unit member of a `Blockchain`. It has a pointer to its parent and holds a collection of `Transaction`s. A `Block` conforms to the `Marshalled` interface which allows it to be serialized and deserialized for network communication.

### Transaction

A `Transaction` is an instruction and holds a set of parameters that a `TransactionHandler` uses to compute the intended update to the state of the system.
