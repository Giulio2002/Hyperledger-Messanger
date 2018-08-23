# Hyperledger-Authenticator

## What
  * this is a simple blockchain made with hyperledger fabric that can provide a simple authentication system(username and password).
  * everyone can use this blockchain to deploy their own authentication system in the network.
  * a peer(called from more users) call the register function to register a new user to the network.
## How to use it
  you have 2 orgs. each org has 2 peers.

  every orgs has the chaincode instantiated.

  every org rapressent a different service(for example amazon account authenticatior nd microsoft account authentication).

  you can enroll new peers or new orgs using fabric-ca.

## How to run it

  ```
  git clone https://github.com/Giulio2002/Hyperledger-Authenticator
  mkdir messanger
  mv Hyperledger-Authenticator/* messanger/
  cd messanger
  sh bootstrap.sh   # it setup,start and give basic rules to the network
  ```
## License

  See LICENSE for more info
