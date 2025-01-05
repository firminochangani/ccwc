/**
 *
 * @param {import("node:fs/promises").FileHandle} f file
 * @param {string} fileName
 * @param {string} option
 * @returns {Promise<string>}
 */
export async function wc(f, fileName, option) {
  const result = {
    totalBytes: 0,
    totalCharacters: 0,
    totalLines: 0,
    totalWords: 0,
  };

  for await (const line of f.readLines()) {
    result.totalBytes += Buffer.from(line).byteLength + 2;
    result.totalLines += 1;
    result.totalWords += line
      .trim()
      .split(/\s+/)
      .filter((word) => word).length;
    result.totalCharacters += line.split("").filter((c) => c).length;
  }

  switch (option) {
    case "-c":
      return `  ${result.totalBytes} ${fileName}`;
    case "-l":
      return `     ${result.totalLines} ${fileName}`;
    case "-w":
      return `   ${result.totalWords} ${fileName}`;
    case "-m":
      return `  ${result.totalCharacters} ${fileName}`;
    default:
      return `    ${result.totalLines}   ${result.totalWords}  ${result.totalBytes} ${fileName}`;
  }
}
