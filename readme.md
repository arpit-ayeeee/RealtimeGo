## Applications

Display realtime data
jprq.io live debugger
k8s API server wartch

# Server sent event:
-> Real-time, unidirectional
-> Pure HTTP requests
-> Supports text only, we have to serialize it into some text format i.e 'data: '
    Maybe when we have to return a pdf of like 64mg, and we have to decode it into
    base64, it can be be costly operation
-> Much less overheard, without using any library, and scaling too
-> Much simpler, native support from browser

# Websockets:
-> Real-time, bidirectional
-> starts with HTTP requests then upgraded to ws, and switches protocol (as it return 101)
-> Supports text as well as binary
-> More overhead, we have to use libraries
-> Complex, as it switches protocol as well as there are some handshakes too
