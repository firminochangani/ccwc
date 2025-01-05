import { open } from "node:fs/promises";
import { wc } from "./wc.js";

async function main() {
  const [, , option, fileName] = process.argv;
  const f = await open(fileName);

  const result = await wc(f, fileName, option);
  console.log(result);

  try {
    await f.close();
  } catch (err) {
    console.error("error closing the file '%s': %s", fileName, err);
    process.exit(1);
  }
}

main();
