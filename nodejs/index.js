const { open } = require("node:fs/promises")

async function main() {
    const result = {
        totalBytes: 0,
        totalCharacters: 0,
        totalLines: 0,
        totalWords: 0
    }

    const [, , option, fileName] = process.argv
    const f = await open(fileName)
    
    for await (const line of f.readLines()) {
        result.totalBytes += Buffer.from(line).byteLength + 2
        result.totalLines += 1

        // FIX ME: Lacks precision
        const segmenter = new Intl.Segmenter([], {granularity: "word"});
        const segmentedText = segmenter.segment(line);
        const words = [...segmentedText].filter(s => s.isWordLike).map(s => s.segment);
        result.totalWords += words.length;

        result.totalCharacters += line.split("").filter(c => c).length;
    }
    
    switch (option) {
        case "-c":
            console.log("  %d %s\n", result.totalBytes, fileName)
            break
        case "-l":
            console.log("    %d %s\n", result.totalLines, fileName)
            break
        case "-w":
            console.log("   %d %s\n", result.totalWords, fileName)
            break
        case "-m":
            console.log("  %d %s\n", result.totalCharacters, fileName)
            break
        default:
            console.log("    %d   %d  %d %s\n", result.totalLines, result.totalWords, result.totalBytes, fileName)
    }

    await f.close().catch(err => console.error("error closing the file '%s': %s", fileName, err))
}

main()

