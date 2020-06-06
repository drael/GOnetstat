# Making GOnetstat faster

I needed a netstat library for some program in go. Problem is the only one that exists is very slow.
So I challenged myself to make it as fast as I can.

I'm adding a branch for each faster version. 
In the readme of each version you can find its execution time.

# Execution time
Original execution time:
```
PASS
ok      GOnetstat       0.396s
```

Current version execution time:
```
PASS
ok      GOnetstat       0.132s
```

# Changes

* Made the main loop that processes the netstat lines use goroutines with channels instead of only using one thread.

# Blog
First off I read through the code and found that the function "netstat" is doing most of the processing.
But it just loops through all the lines, and each line takes some time to process. Go routines helped that a lot.

