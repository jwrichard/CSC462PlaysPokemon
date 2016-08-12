# CSC462PlaysPokemon
Welcome to Twitch Plays Pokemon - Distributed Systems Version!

Created twitch plays pokemon, but not hosted on twitch. Where does distributed work take place? 
- Playing under the rules of democracy...MAP REDUCE!
- Create web front end that talks to master node
- Utilize AWS resources to execute workload (EC2 & EFS)
- Master node divides work up to slaves to be parsed
- Slaves return a total of each move entered into chat to be aggregated
- Master sends the chosen move decided on by Map Reduce to the host, and the game executes the move

# Screenshot
![Alt text](example.jpg?raw=true "Sample gameplay")

# Key notes
- Nodes require multiple extra files/symlinks which could not be included, therefore you cannot simply download and redeploy this
- Nodes are no longer active (I don't feel like paying to keep it running)

# Nodes
## Master
- __IP__: 54.213.250.201
- __RTMP Server__: rtmp://54.213.250.201/live (Stream key: "test")
- __User__: ubuntu 

## Slave 1
- __IP__: 54.201.255.6
- __User__: ubuntu 

## Slave 2
- __IP__: 54.208.103.21
- __User__: ubuntu 

