# Making GOnetstat faster

I needed a netstat library for some program in go. Problem is the only one that exists is very slow.
So I challenged myself to make it as fast as I can.

I'm adding a branch for each faster version. 
In the readme of each version you can find its execution time.

# Execution time
Original execution time:
```
PASS
ok      GOnetstat       0.955s
```

Current version execution time:
```
PASS
ok      GOnetstat       0.079s
```

# Changes

* Added a buffer to the channels which can free resources (less goroutines blocking)

# Blog
I read a lot about channels then I came to the conclusion that I needed a buffer on my channels. A lot of goroutines blocked creates a lot of OS threads. Freeing them can make the program use less resources.

