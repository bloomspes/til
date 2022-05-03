const { log } = console;

// Create a stream object.
// The `stream` module provides an API for implementing the stream interface.

const { Stream } = require("stream");
const writableStream = new Stream.Writable()
const readableStream =  new Stream.Readable({
    read() {}
})

// Implement _write.
writableStream._write = (chunk, encoding, next) => {
    log(chunk.toString())
    next()
}

// Reading a data from a readable stream.
readableStream.pipe(writableStream)

readableStream.push('hi')
readableStream.push('file system')

readableStream.on('close', () => writableStream.end())
writableStream.on('close', (stream) => log('ended'))

// Destroy the stream. 
// Implementors should not override this method, but instead implement `readable._destroy()`.
readableStream.destroy()