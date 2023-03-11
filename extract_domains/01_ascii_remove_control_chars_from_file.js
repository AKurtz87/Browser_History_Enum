const fs = require("fs");

const rows = 2;
const columns = 97;
const characters = ["#", "@", "%", "A", "O", "X"];
const colors = ["red", "green", "blue", "yellow", "cyan", "magenta"];
let image = "";

for (let row = 0; row < rows; row++) {
  for (let col = 0; col < columns; col++) {
    const character = characters[Math.floor(Math.random() * characters.length)];
    const color = colors[Math.floor(Math.random() * colors.length)];
    image += `\x1b[38;5;${colors.indexOf(color)}m${character}\x1b[0m`;
  }
  image += "\n";
}

function bufferToAscii(buffer, charsToReplace) {
  let ascii = buffer.toString("ascii");
  charsToReplace.forEach((char) => {
    ascii = ascii.split(char).join(" ");
  });
  return ascii;
}

async function convertFileToAscii(
  inputFilePath,
  outputFilePath,
  charsToReplace
) {
  console.log("\n" + image);
  try {
    const data = await fs.promises.readFile(inputFilePath);
    const ascii = bufferToAscii(data, charsToReplace);
    const output = ascii.match(/.{1,50}/g).join("\n"); // max 50 characters per line
    await fs.promises.writeFile(outputFilePath, output);
    console.log(
      "                                     Conversion complete!\n\n" + image
    );
  } catch (error) {
    console.error(error);
  }
}

const charsToReplace = [
  "\x00", // null character
  "\x01", // start of heading
  "\x02", // start of text
  "\x03", // end of text
  "\x04", // end of transmission
  "\x05", // enquiry
  "\x06", // acknowledge
  "\x07", // bell
  "\x08", // backspace
  "\x09", // horizontal tab
  "\x0A", // line feed
  "\x0B", // vertical tab
  "\x0C", // form feed
  "\x0D", // carriage return
  "\x0E", // shift out
  "\x0F", // shift in
  "\x10", // data link escape
  "\x11", // device control 1
  "\x12", // device control 2
  "\x13", // device control 3
  "\x14", // device control 4
  "\x15", // negative acknowledgement
  "\x16", // synchronous idle
  "\x17", // end of transmission block
  "\x18", // cancel
  "\x19", // end of medium
  "\x1A", // substitute
  "\x1B", // escape
  "\x1C", // file separator
  "\x1D", // group separator
  "\x1E", // record separator
  "\x1F", // unit separator
  "\x7F",
];

const inputFilePath = process.argv[2];
const outputFilePath = "asciiFile.txt";

convertFileToAscii(inputFilePath, outputFilePath, charsToReplace);
