# Making GOnetstat faster

I needed a netstat library for some program in go. Problem is the only one that exists is very slow.
So I challenged myself to make it as fast as I can.

I'm adding a branch for each faster version. 
In the readme of each version you can find its execution time.

# Execution time
Original execution time:
```
PASS
ok      GOnetstat       0.446s
```

Current version execution time:
```
PASS
ok      GOnetstat       0.076s
```

# Changes

* Made the findpid function only get all process file descriptors once instead of for every netstat line.

# Blog
This time I used go profiling and found that the main bottleneck is the findpid function.
So I made it run the function once and pass a pointer of the result to every goroutine processing the netstat lines.


