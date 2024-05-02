# Blockchain in Go


#### 4 STEPS TO CREATING A BLOCKCHAIN
1. Create a block.
2. Add the data (header and body) to the block.
3. Hash the block.
4. Chain the blocks together.

#### What Prerequisites Do I Need to Create a Blockchain?
* A network.
* Cryptography.
* A data structure and algorithms.
* Decentralized systems.
* Javascript, Go or Python knowledge.

*You only need to understand the basic concepts to program your first blockchain prototype, so let’s begin with some theories.*


## What Is a Blockchain?

* **You can’t program something you don’t understand, right? Let’s break it down. Blockchain is not Bitcoin. Blockchain is not a digital currency, Blockchain is a set of different technologies that had already existed before its creation.**

* Let’s simplify things with an example.

* Let’s take a MySQL database that stores some information in different tables.
*With the above database, there are some limitations. It allows you to do some tasks, including:*

   1. **Create, retrieve, update and delete (CRUD) operations.**
   2. **Storing the same information twice.**

*But its limitations are many, including:*

   1. **We can drop the entire database either mistakenly or intentionally.**
   2. **We can’t share sensitive information with others.**
   3. **It’s centralized, which means there’s a single point of failure and a security issue.**
   4. **There’s no way to trust everything that is stored in it.**
   5. **Ill-intentioned people can comprise the database.**
   6. **The database needs a manager.**
   7. **The users don’t have power over their own data**


## How Does Blockchain Work?

* **In other words, a blockchain is a chain of blocks. These blocks are like tables in the database, but they can’t be deleted or updated. The blocks contain information such as transactions, nonce, target bit, difficulty, timestamp, previous block hash, current block hash, Markle tree and block ID, etc, and the blocks are cryptographically verified and chained up to form an immutable chain called a blockchain or a ledger.**

* **The same chain is then distributed to all the nodes (computers or miners) across the network via a P2P network. So, instead of a centralized database, all the transactions (data) that are shared across the nodes are contained in blocks, which are chained together to create the ledger. This ledger represents all the data in the blockchain. All the data in the ledger is secured by cryptographic hashing and digital signature and validated by a consensus algorithm. Nodes on the network participate to ensure that all copies of the data distributed across the network are the same**


#### 5 KEY CONCEPTS IN THE BLOCKCHAIN ECOSYSTEM
1. Cryptographic hash and digital signature.
2. Immutable ledger.
3. P2P network.
4. Consensus algorithm ( PoW, PoS, PBFT, ETc…).
5. Block validation ( Mining, Forging, etc…).


## What Are the Benefits of Using Blockchain?
* **There are a number of benefits to using blockchain, including:**

1. Removing intermediary organizations.
2. An immutable ledger
3. Transparency
4. Security
5. Reliability


## When Should You Use Blockchain?
* **Blockchain is not a silver bullet, so, it’s best to use it when:**

1. The data stored cannot be modified (proof of existence ).
2. The data cannot be denied by its owner (non-repudiation).
3. You want decentralization.
4. You want one source of truth.
5. You want high security.
6. You don’t care about speed. For example, Bitcoin takes 10 minutes on average to validate a block of transactions. But some blockchains are faster because they use different consensus algorithms other than PoW. We will talk about this later.

## Blockchain Use Cases
* **Blockchain goes beyond cryptocurrency and bitcoin. It can be used in different sectors, including:**

1. Real estate: Land ownership.
2. Healthcare: Securely record the patient’s data.
3. Finance: Reduce the taxes and intermediaries, anti-money laundering and cross border payment.
4. Supply chain: Track items from the vendors to the customers, including verifying authenticity and original content creation.
5. Cybersecurity: DDOS attacks.
6. Giving the power back to the user: Own your data and share it securely with whom you want (DID).
7. Cryptocurrency.
8. Voting mechanism.

## Blockchain Platforms and Applications to Know
* Bitcoin
* Ethereum
* Hyperledger Fabric
* EOS
* Chainlink
* Cardano
* Dogecoin (meme coin)


## Blockchain Types
* **There are three types of blockchain:**

* **Private:** Use only internally, and when we know the users (eg. Hyperledger Fabric).
* **Public:** Everyone can see what is going on (Bitcoin, Ethereum).
* **Hybrid:** In case you want to combine the first two.


## Create Your Own Blockchain 
* There are two ways to build a blockchain

1. The easiest way is to use a pre-built blockchain open-source like Ethereum (create distributed applications, altcoins, decentralized finance (DeFi) and non-fungible tokens (NFTs), etc.), Fabric (configure a private blockchain), EOS or Cardano, etc. so that you don’t have to deal with the core engine, which is difficult to implement.

2. If that doesn’t fit your requirements, then you can build one from scratch or fork, modify and/or improve an existing blockchain open-source code. For example, Litecoin and Bitcoin cash were forked from Bitcoin. This last method is tougher, more time-consuming and requires a lot of work and a strong team.