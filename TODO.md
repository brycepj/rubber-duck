- architecture should be:
  - browser app (wraps wasm webrtc client)
  - server app (wraps go webrtc client) 
  - both apps act as a client that facilitate the exchange. the webrtc clients will implement features of webrtc, that we can build on top of 

- http web server that returns an html page that fetches a wasm file, compiles and runs it.
  - simple go server that you can run locally. It should have a rest API that returns an html file
  - Add JS to the html file that fetches wasm and runs it. 
  - Add an endpoint to the go server that returns the wasm
  - Probably use some CRA to do something that actually works and runs and you can extend easily
  - Get a hello world with the webassembly module running and logging something in the browser

- Go from there...
