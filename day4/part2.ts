import fs from "fs";

const input = fs.readFileSync("input.txt", "utf8").trim().split('\n').map((row) => row.split(''));

let count = 0;

for (let i = 0; i < input.length; i++) {
  for (let j = 0; j < input[i].length; j++) {
    const current = input[i][j];

    if (current === "A") {

      /**
        *  M . M
        *  . A .
        *  S . S
        */

      if (
        input[i - 1]?.[j - 1] === "M" &&
        input[i + 1]?.[j + 1] === "S" &&
        input[i - 1]?.[j + 1] === "M" &&
        input[i + 1]?.[j - 1] === "S"
      ) {
        count++;
        continue;
      }


      /**
        *  S . S
        *  . A .
        *  M . M
        */

      if (
        input[i - 1]?.[j - 1] === "S" &&
        input[i + 1]?.[j + 1] === "M" &&
        input[i - 1]?.[j + 1] === "S" &&
        input[i + 1]?.[j - 1] === "M"
      ) {
        count++;
        continue;
      }

      /**
      * M . S
      * . A .
      * M . S
      */

      if (
        input[i - 1]?.[j - 1] === "M" &&
        input[i + 1]?.[j + 1] === "S" &&
        input[i + 1]?.[j - 1] === "M" &&
        input[i - 1]?.[j + 1] === "S"
      ) {
        count++;
        continue;
      }

      /**
      * S . M
      * . A .
      * S . M
      */

      if (
        input[i - 1]?.[j - 1] === "S" &&
        input[i + 1]?.[j + 1] === "M" &&
        input[i + 1]?.[j - 1] === "S" &&
        input[i - 1]?.[j + 1] === "M"
      ) {
        count++;
      }

    }
  }
}

console.log(count);
