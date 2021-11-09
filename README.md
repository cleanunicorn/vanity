# Vanity

Bruteforce private keys that have the associated public key with a matching prefix/suffix.

## Run 

Start the program, it will ask for a prefix or suffix, they are both optional; you can search for a prefix or a suffix or both.

The number of threads defaults to your CPU core count.

```sh
$ go run .
Prefix: 0xc001
Suffix: d00d
Number of threads: [24] 
Using 24 threads to search for an address that looks like 0xc001...d00d 
Speed: 200000 keys/sec, Total: 200000
Speed: 300000 keys/sec, Total: 300000
Speed: 200000 keys/sec, Total: 400000
Speed: 250000 keys/sec, Total: 500000
Speed: 200000 keys/sec, Total: 600000
Speed: 233333 keys/sec, Total: 700000
...
```

After it finds an address, it will output the (checksummed address)[https://github.com/ethereum/EIPs/blob/master/EIPS/eip-55.md] and the associated private key.
