# Proof of Work

Proof of Work algorithms must meet a requirement doing the work is hard, but verifying the proof is easy.

## Hashing

A hash is a unique representation of the data it was calculated on. A hash function is  a function that takes data of arbitrary size and produces a fixed size hash.

- Original data cannot be restored from a hash. Thus, hashing is not encryption
- Certain data can have only one hash and the hash is unique
- Changing even one byte in the input data will result in a completely different hash

In blockchain, hashing is used to gurantee the consistency of a block. The input data for a hashing algorithm contains the hash of the previous block, thus making it impossible to modify a block in the chain: one has to recalculate its hash and hashes of all blocks after it.

## Hashcash

- Take some publicly know data (in case of Bitcoin, it's block headers)
- Add a counter to it. The counter starts at 0.
- Get a hash of the data + counter function.
- Check that the hash meets certain requirements.
1. If it does, you're done
2. If it doesn't, increase the counter and repeat the steps 3 and 4

Thus, this is brute force algorithm: you change the counter, caculate a new hash, check it, increment the counter, calculate a hash, etc. That's why it's computationally expensive. 

## Implementation