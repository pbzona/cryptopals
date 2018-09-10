# Single-byte XOR cipher

The hex encoded string:

`1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736`

... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.

**Solution** - No test written for this one because there's no given result to look for. My solution prints out scored output - correct answer is actually the second top score, but I believe this has more to do with the frequency table than my actual method.

**Update** - Playing with some different frequency tables, I found that you can get the correct result with a much simpler scoring system - assign each key a linear number value with 1 being the least frequent and 26 being the most frequent. I've removed the [previous table I was using](http://pi.math.cornell.edu/~mec/2003-2004/cryptography/subs/frequencies.html) in light of this, and made a mental note not to overthink other challenges in the first tier 😎