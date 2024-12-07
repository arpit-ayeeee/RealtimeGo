## Server sent event:
-> Real-time, unidirectional
-> Pure HTTP requests
-> Supports text only, we have to serialize it into some text format i.e 'data: '
    Maybe when we have to return a pdf of like 64mg, and we have to decode it into
    base64, it can be be costly operation
-> Much less overheard, without using any library, and scaling too
-> Much simpler, native support from browser

# Applications
-> Display realtime data
-> jprq.io live debugger
-> k8s API server wartch

## Websockets:
-> Real-time, bidirectional
-> starts with HTTP requests then upgraded to ws, and switches protocol (as it return 101)
-> Supports text as well as binary
-> More overhead, we have to use libraries
-> Complex, as it switches protocol as well as there are some handshakes too
-> WS is purely on top of TCP

# Applications
-> Chat apps
-> Collaborative apps
-> real-time updates
    - stock ticks (sse can be used too)
    - football match (sse can be used too)
    - election results (sse can be used too)


## Web RTC (Real real-time apps)
-> Works peer to peer, eg. video call, audio call
-> Here we have stun servers which help two devices connect peer to peer
