# pktdump

A simple network packet decoder in Go

```
$ cat packet.txt
0022 4466 88aa 1133 5577 99bb 0800 4500
004e 1779 4000 4006 3c0a 0a00 0001 0808
0808 8707 09d2 82f2 cfda f065 feb4 8018
00e5 71e1 0000 0101 080a 6d8f 9a46 5d5b
f136 6161 6161 6161 6161 6161 6161 6161
6161 6161 6161 6161 6161 6161

$ cat packet.txt | xxd -r -p | ./pktdump
== Begin Ethernet ==
MAC (dst): 00:22:44:66:88:aa
MAC (src): 11:33:55:77:99:bb
Ethertype: 0x0800
  == Begin IPv4 ==
  TTL: 64
  IP Address (src): 10.0.0.1
  IP Address (dst): 8.8.8.8
    == Begin TCP ==
    Port (src): 34567
    Port (dst): 2514
      == Begin Raw ==
      aaaaaaaaaaaaaaaaaaaaaaaaaa
      == End Raw ==
```
