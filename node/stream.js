const { log } = console;

// create a stream object
const { Stream } = require("stream");
const writableStream = new Stream.Writable()
const readableStream =  new Stream.Readable({
    read() {}
})

// implement _write
writableStream._write = (chunk, encoding, next) => {
    log(chunk.toString())
    next()
}

// reading a data from a readable stream
readableStream.pipe(writableStream)

readableStream.push('hi')
readableStream.push('file system')

readableStream.on('close', () => writableStream.end())
writableStream.on('close', (stream) => log('ended'))

readableStream.destroy()