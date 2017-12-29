# GLORIOUS GLORY OF PACKET SLANGING

This just passes butter. build it, run it. 

Imagine, if you will, you built it with the name `passer`. You would then run it like this:

`./passer -in 8080 -out 80`

This will expose whatever is already running internally on port TCP80, on port TCP8080. Say you have a webserver tied to 127.0.0.1:80 and want to just arbitrarily expose it to other machines on your network on port 8080. Run it like I've written above. 

-in: this programs listens on this TCP port

-out: this program sends traffic off to something running on this port

Hopefully this is crazy clear. It's not like this diddy is long, or responsibly-written. Heck, I didn't even reuse my validator from another thing I wrote, so don't go passing it port POLARBEAR or some such nonsense.

-the wicked jsmonet
