# Making GOnetstat faster

I needed a netstat library for some program in go. Problem is the only one that exists is very slow.
So I challenged myself to make it as fast as I can.

I'm adding a branch for each faster version. 
In the readme of each version you can find its execution time.

# Execution time
Original execution time:
```
PASS
ok      GOnetstat       0.542s
```

Current version execution time:
```
PASS
ok      GOnetstat       0.049s
```

# Changes

* First reads all process links then iterates over it instead of reading all links in each of the netstat lines

# Blog
After profiling I saw that readlink is taking most of the cpu time. After some thinking I got to the conclusion that its because each line is reading all process links.
So I'm reading all the links at the start.


