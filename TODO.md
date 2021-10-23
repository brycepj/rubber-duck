- architecture should be:
  - browser app (wraps wasm webrtc client)
  - server app (wraps go webrtc client) 
  - both apps act as a client that facilitate the exchange. the webrtc clients will implement features of webrtc, that we can build on top of 


// this is up to date for now 

- http web server that returns an html page that fetches a wasm file, compiles and runs it.
  - Add wasm request handler to go server
  - Add JS to the html file that fetches wasm and runs it. 
  - Get a hello world with the webassembly module running and logging something in the browser

- Go from there...
