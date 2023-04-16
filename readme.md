# WolServer
#### a server that makes it safe to use wol from the network

### usage
___

You can skip the first 3 steps. To do this, just use the `wol.exe` file I pre-compiled.
In this case, just start from step 4

1. Make sure you have golang installed.
2. Clone or download the repository:

```git clone  https://github.com/utherbit/WolServer```

3. build the project.

```go build```

4. Create a file containing the computer's mac address and wake key.

by default the file should be named input.txt but this can be changed in the .env file
example:
```
<wake-up-key> <mac-address> // comment
2b55d960-ba53-415c-ab25-fa197ba3977a A0-B1-C2-D3-E4-F5 // utherbit pc
8a96f8ec-be9c-4e55-98ed-699f64a0d8e9 B2-A3-E5-F8-A0-C1 // example user2 
```
5. Run the compiled project on your server.

make sure that next to the executable file (in the same directory) are `.env` and `input.txt` files


as a result, you should see the following inscription 

```
Listen <host>:<port>
```


This means you can test the server, 
for this go to the page 

`http://<host>:<port>/wakeup/<token>` 

after changing the `host`, `port`, `token` 
values. 

token - the key given in the `input.txt` file for the wakeup PC


On the page you should find the following entry:
Package sent successfully.


this means that the request was successfully processed by the server and the wol packet was sent