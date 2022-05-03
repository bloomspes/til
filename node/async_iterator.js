
const fs = require('fs');
const { log } = require('console');

const { promisify } = require('util');
const { once } = require('events');

// Initialize stream object.
const Stream = require('stream');
const readableStream = new Stream.Readable()


async function sample(readable) {
    for await (const chunk of readable) {
        log(chunk);
    }
}

const readable = fs.createReadStream('node/test.txt', 'utf-8')
sample(readable);

// Default version of stream.finished() is callback-based,
// but can be turned into a Promise-based version.
const finished = promisify(Stream.finished);

async function writeFile(iterable, filePath) {
    const writable = fs.createWriteStream(filePath, 'utf-8');
    for await (const chunk of iterable) {
        if (!writable.write(chunk)) {
            await once(writable, 'drain');
        }
    }
    
    // Closing a writable stream.
    // no more data will be written to the Writable.
    writable.end();

    await finished(writable);
}

writeFile(['one\n', 'hello world'], 'node/sample.txt');
