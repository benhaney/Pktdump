# pktdump

A simple network packet decoder in Go

```
$ cat examples/eth-ip4-tcp.hex
0011 228b 9411 5533 aa0c a508 0800 4500
0058 63dd 4000 4006 499e 0a00 0001 acd9
0c4e 8a10 0050 8e67 c702 8b05 fe8d 8018
00e5 57ef 0000 0101 080a 58c2 a8da 136c
4f75 4745 5420 2f20 4854 5450 2f31 2e31
0d0a 486f 7374 3a20 676f 6f67 6c65 2e63
6f6d 0d0a 0d0a

$ cat examples/eth-ip4-tcp.hex | xxd -r -p | ./pktdump
== Begin Ethernet ==
MAC (dst): 00:11:22:8b:94:11
MAC (src): 55:33:aa:0c:a5:08
Ethertype: 0x0800
  == Begin IPv4 ==
  TTL: 64
  IP Address (src): 10.0.0.1
  IP Address (dst): 172.217.12.78
    == Begin TCP ==
    Port (src): 35344
    Port (dst): 80
      == Begin Raw ==

GET / HTTP/1.1
Host: google.com

      == End Raw ==

$ cat examples/eth-ip6-udp.hex
0017 108b 9411 503e aa0c a508 86dd 6007
d796 0015 1140 2605 6000 3fc0 0000 0000
0000 0000 2514 2607 f8b0 4000 080c 0000
0000 0000 200e 9d06 0050 0015 0e14 6161
6161 6161 6161 6161 6161 61

$ cat examples/eth-ip6-udp.hex | xxd -r -p | ./pktdump
== Begin Ethernet ==
MAC (dst): 00:17:10:8b:94:11
MAC (src): 50:3e:aa:0c:a5:08
Ethertype: 0x86dd
  == Begin IPv4 ==
  Hop Limit: 64
  IP Address (src): 2605:6000:3fc0:0000:0000:0000:0000:2514
  IP Address (dst): 2607:f8b0:4000:080c:0000:0000:0000:200e
    == Begin UDP ==
    Port (src): 40198
    Port (dst): 80
      == Begin Raw ==

aaaaaaaaaaaaa

      == End Raw ==
```
