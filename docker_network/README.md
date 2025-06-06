# _DOCKER NETWORKING_

## _Theory Part_

Basically, each container may have to communicate with either the host or another container via network IP only. Sometimes, the container should be logically isolated so that it won't communicate with other containers.

In case of VMs, they are isolated at the hardware level by the hypervisor, each with its own operating system and resources. Communication between VMs usually happens over a virtual network and requires them to be in the same subnet or to have proper routing rules and firewall configurations if they are in different subnets or VPCs.


***<ins>Bridge Network</ins>***: Usually there shouldn't be a network communication between the host & containers due to the differnce in their subnets, hence docker came up with ***<ins>`Bridge Network`</ins>*** concept where docker creates a virtual network `veth0 (docker0)` and can communicate with host/container's `eth0`. If there are any unnecessary changes to the Bridge Network config or if anyone deletes it, the containers would never talk to the host. Without it nullifies the purpose of your application inside the container serving the user will be void. 

In case of logically isolated containers, out-of-the-box Bridge Network is not feasible as it won't provide any security and is vulnerable for hacker attacks. For this case, we can create our own ***<ins>`Custom Network`</ins>*** using ***<ins>`Docker Bridge Network`</ins>*** command, which will isolate the container(s) from others where it won't interact with veth0 (docker0).

## _Practical Part_

Step 1: Create a first container using the following command  
```
CMD -> docker run -d --name <CONTAINER_NAME> <IMAGE_NAME>:<TAG>
EX  -> docker run -d --name login nginx:latest
```
Ubuntu and Alpine are both minimal base images, meaning they don’t have a default process running that keeps the container alive. That’s why, when you run them in Docker, they immediately exit unless you provide a persistent command such as `bash -c "while true; do sleep 30; done"` .
```
CMD -> docker run -d --name <CONTAINER_NAME> <IMAGE_NAME>:<TAG> <COMMAND>
EX  -> docker run -d --name login ubuntu:latest bash -c "while true; do sleep 30; done"
```
Step 2: Create a second container using the following command
```
CMD -> docker run -d --name <CONTAINER_NAME> <IMAGE_NAME>:<TAG>
EX  -> docker run -d --name logout nginx:latest
```
or
```
CMD -> docker run -d --name <CONTAINER_NAME> <IMAGE_NAME>:<TAG> <COMMAND>
EX  -> docker run -d --name logout ubuntu:latest bash -c "while true; do sleep 30; done"
```
Step 3: Now, try to get the IP address of the containers that we have created. Even though the containers are created from the host, they would have serial CIDR subnet IP's such as - 172.17.0.2, 172.17.0.3, 172.17.0.4, so on and so forth. Still, it's always a good practice to use the following command to get the IP address in case we have too many containers. 

> [!NOTE]
> using `docker inspect` along with `grep -i` & `grep -E` to ignore case sensitive and enabling extended regular expressions. 

```
CMD -> docker inspect <CONTAINER_NAME1> <CONTAINER_NAME2> <CONTAINER_NAME3> | grep -Ei '"ipaddress":|"networkmode":'
EX  -> docker inspect login logout upi_payment | grep -Ei '"ipaddress":|"networkmode":'
```
or
```
CMD -> docker inspect <CONTAINER_NAME1> <CONTAINER_NAME2> <CONTAINER_NAME3> | grep -E '"IPAddress":|"NetworkMode":'
EX  -> docker inspect login logout upi_payment | grep -E '"IPAddress":|"NetworkMode":'
```

> [!NOTE]
> Similarly, we can use the jq command, which is often mistaken for "JSON Query," but it stands for "JSON Processor."

```
CMD -> docker inspect <CONTAINER_NAME1> <CONTAINER_NAME2> <CONTAINER_NAME3> | jq '.[]|.NetworkSettings|.Networks'
EX  -> docker inspect login logout upi_payment | jq '.[]|.NetworkSettings|.Networks'
```
Step 4: Now, enter any of the containers that we have created to see if we can ping other containers via `veth0 (docker0) Bridge Network`
```
CMD -> docker exec -it <CONTAINER_NAME> bash
EX  -> docker exec -it login bash
```
or
```
CMD -> docker exec -it <CONTAINER_NAME> /bin/bash
EX  -> docker exec -it login /bin/bash
```
Step 5: Upon entering into the container, run the commands to ping the other container, as out-of-the-box containers won't have ping installed.
```
CMD -> apt update && apt-get install iputils-ping -y
```
Step 6: Upon installation, try to ping the Bridge Network or other containers with their IP address. `-c` flag to limit the count of pings.
```
CMD -> ping -c 1 <IP_ADDRESS>
EX  -> ping -c 1 172.17.0.2
```	
Now, you will see the ping response as the host and containers are in the same subnet because of Docker Bridge Network out of the box.

Step 7: If we want to list out all the Docker Bridge Networks on the host
```
CMD -> docker network ls
```	
Step 8: If we want to create a custom Docker Bridge Network, use the following command
```
CMD -> docker network create <NETOWRK_NAME> 
EX  -> docker network create secure-network
```
Step 9: If we want to remove a custom Docker Bridge Network, use the following command
```
CMD -> docker network rm <NETOWRK_NAME> 
EX  -> docker network rm secure-network
```
Step 10: Now, using the below command, we will assign the newly created network to a new container, which will be logically isolated from other container,s where we can run secure applications
```
CMD -> docker run -it --name <CONTAINER_NAME> --network <NETOWRK_NAME> <IMAGE_NAME>:<TAG> <COMMAND>
EX  -> docker run -it --name upi_payment --network secure-network ubuntu:latest bash -c "while true; do sleep 30; done"
```
or
```
CMD -> docker run -it --name <CONTAINER_NAME> --network <NETOWRK_NAME> <IMAGE_NAME>:<TAG>
EX  -> docker run -it --name upi_payment --network secure-network nginx:latest
```
Step 11: Now, using the below command, we will assign the host network to a new container and see what happens
```
CMD -> docker run -it --name <CONTAINER_NAME> --network <NETOWRK_NAME> <IMAGE_NAME>:<TAG> <COMMAND>
EX  -> docker run -it --name host --network host ubuntu:latest bash -c "while true; do sleep 30; done"
```
or
```
CMD -> docker run -it --name <CONTAINER_NAME> --network <NETOWRK_NAME> <IMAGE_NAME>:<TAG>
EX  -> docker run -it --name host --network host nginx:latest
```

> [!IMPORTANT]
> If you inspect the newly created host container, you won't see any IPAddress as the host is already bound with the host network itself, hence Docker didn't create any virtual network in this case.

Step 12: Now, if you try to go inside any of the login/logout containers and post installation of iputils-ping, we won't be able to ping the upi_payment container as it will be created in a different subnet and vice-versa for upi_payment.

> [!IMPORTANT]
> In this way, we can achieve communication and logically isolated containers using Docker Default Bridge & Custom Network Bridge.

![image](https://github.com/user-attachments/assets/b7558575-66d4-4c6a-89f3-78227cd6a884)


