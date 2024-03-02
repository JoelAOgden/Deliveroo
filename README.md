# Read me

I'm not too familiar with cron (which now I know about it i'm surprised by). I suspect
there might be a few misunderstandings about how cron is implemented.

I didn't have much time over the weekend, been pretty busy so had to do this last minute on a saturday. 
I managed to get all the cron bits working and set up a simple service to use them.

Sadly there wasn't much time to turn this into a cmd application.

Not something I've done much in golang, so didn't prioritise it. I'm sure it's nothing difficult though.

Overall I've tried to focus more on the readability of things, I had a look at some other cron libraries
and I couldn't find a very "clean" way to parse it. Hopefully the approach I've taken has seperated out 
some of the gross bits into digestible functions.

### known issues
  - Ordering 
    - simple fix, just add sorting at the end
  - Duplicates 
    - easy to fix with a map or something
  - super inefficient 
    - all the nested function's eat the heap, pass by val also is also crazy memory inefficient here
    - I dunno, i'm sure there's a much more effective way that what i'm doing, but n<1000 so not worth worrying about


### Improvements
  - Testing is pretty meh
  - Documentation is pretty terrible :(
  - my brain keeps wanting to add an interface for a parser but I can't think of why (shrug)
  - go routines would be nice and bloody easy with how it's structured
    - but I highly doubt it'll do much considering the simplicity probably isn't worth the locking
  