# Making GOnetstat faster

I needed a netstat library for some program in go. Problem is the only one that exists is very slow.
So I challenged myself to make it as fast as I can.

I'm adding a branch for each faster version. 
In the readme of each version you can find its execution time.

# Execution time
Original execution time:
```
PASS
ok      GOnetstat       0.413s
```

Current version execution time:
```
PASS
ok      GOnetstat       0.045s
```

# Changes

* Made the process array and the inodes array statically lengthed

# Blog
After a bit of profiling I saw runtime.futex using a most of the cpu time.
So I made all the process and inodes array statically lengthed, which removes the need for the append function.


