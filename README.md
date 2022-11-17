# vesting_program
After unzipping the file, and going to vesting_program directory, please run the 
following command to install the dependencies: 
1. From the subdirectory, run the following command to compile the program:
```go build ```
2. If you want to run the tests please following command:
``` go test ./workers```


I decided to create two workers to handle the vesting process. The first worker is responsible 
for reading the csv file and creating the vesting accounts. The second worker is responsible 
for vesting the accounts. The second worker will be triggered by the first worker. The first worker will create 
a channel and send the channel to the second worker. The second worker will read from the channel and vest the accounts. 
The second worker will also send a message to the first worker when it is done vesting the accounts. 
The first worker will then close the channel and exit the program.

I have 12 cores on my machine, so I decided to create 12 workers for the second worker. But it will depend on number of 
cores on YOUR machine. I divide number of lines from the csv file by 12 to get the number of lines each worker will handle.

I also decided to use a buffered channel to send the lines from the first worker to the second worker. I decided to use a
buffered channel because I don't want the first worker to wait for the second worker to read from the channel. I want the
first worker to keep reading from the csv file and send the lines to the second worker. 

