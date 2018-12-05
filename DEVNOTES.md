# Development Notes

[Home](README.md) | [Installation](INSTALLATION.md) | [Development Notes](DEVNOTES.md) | [Licence](LICENSE)

### 2018-12-04 Clicktype Enumeration Values + Logging
Added some code to log the types of button clicks that can occur.  While the [Button Guide](https://docs.aws.amazon.com/iot-1-click/latest/developerguide/button-guide.html)
 documents the types, it doesn't indicate enumeration of possible values.

### 2018-12-04 Tesla Wakeup Command Responsiveness
Taking a look at the logging for the Wakeup API it appears on average that the action is taking in excess
of 17 seconds at 95th percentile.  I filed bug #12, to extend the retry time and make that configurable.

### 1) Magic names for Golang binaries
The recipe to get started with Golang and Lambda is superficially a little messy and took some trial and error to get right.
First the error messages are misleading - ```no such file or directory```.

Turns out that Lambda expects the binary name, and the function name 
for your lambda entrypoint to be the same.  30mins of my
life I'll never get back.

### 2) Golang IotButton Samples
Moreover there are no official samples for IotButtonEvent written in Go at this point that appear in Google.
If they do exist, they are buried.  The events.IoTButtonEvent message types recently committed to the 
mainline of the AWS Go SDK but unfortunately not labelled in a release.  Not a show stopper, but made ```dep``` more 
messy. A little prodding reminded the owners to move the release forward to 1.7.0.

### 3) Security and Lambda
I'll pass credentials for my Tesla API key into the lambda through environment variables. Investigate how to securely 
attach those to the Lambda definition.  Amazon Key Management API enables creation of secrets, and a client-side API to 
decode secrets within the code.  Cipher key is exchanged externally as part of the
IAM role/context of the Lambda

### 4) Tesla OAuth Secrets
Tesla's API Gateway has a set of OAuth "Password Grant" secrets to grant access to the token exchange API.  
Cadence of change as yet unknown. They were initially extracted from the 
Android application approximately 1yr ago. 
(notes from Reddit here: https://www.reddit.com/r/teslamotors/comments/72ilu3/a_few_tidbits_from_disassembling_the_tesla/)
I'd imagine that one-day we'll need some sort of service token or MFA to generate these secrets.

Some working secrets appear to be here - https://pastebin.com/pS7Z6yyP

### 5) PR: ElectricGopher Tesla API Libraries
Looking around for Tesla API libraries and instructions for the Tesla Owners REST API. [Tim Dorr's Tesla Ruby API](https://github.com/timdorr/tesla-api) is conceptually complete
and has a great cassette reference of the interactions with the Tesla API server.  I used Tim's code in conjunction with [Nathan Beyer's](https://nbeyer.io) partial Golang port [ElectricGopher](https://github.com/nbeyer/electricgopher) to 
port to get started on a functional client - currently working in ```vehicle-fix``` branch - [https://github.com/stephbu/electricgopher/]()

Some things I've fixed/added so far: 
- BUG:Socket leakage from dangling response stream
- HTTP POST client support
- Vehicle Command Library
     - Wakeup
     - Unlock
     
### 6) BUG: Containment, Package Init, and Lambda
Spent a bunch of time refactoring the pkg structure, moving handlers out of ```main``` in ```pkg```.  In doing so
seemed to run into a bug where static construction through package init() was no-longer being called.
Worked fine in the HandlerTest harness, failed with nil-pointer assertions in Lambda host.
