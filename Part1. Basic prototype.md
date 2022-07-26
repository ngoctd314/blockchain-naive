# Basic prototype

## Introduction

Blockchain is unique because it's not a private database, but a public, i.e everyone who uses it has a full or partial copy of it. And a new record can be added only with a consent of other keepers of the databases.

## Block

In blockchain it's block that store valueable information. For example, bitcoin blocks store transactions, the essence of any cryptocurrency. Besides this, a block contains some technical information, like its version, current timestamp and the hash of the previous block.

## Blockchain

In its essence blockchain is just a database with certain structure: it's an ordered, back linked-list. Which means that blocks are stored in the insertion order and each block is linked to the previous one. This structure allows to quickly get the lastest block in a chain and to (efficiently) get a block by its hash.
