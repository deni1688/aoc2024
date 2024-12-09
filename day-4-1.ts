import fs from "fs";

const input = fs.readFileSync("input.txt", "utf8").trim().split('\n').map((row) => row.split(''));

let count = 0;

for (let i = 0; i < input.length; i++) {
  for (let j = 0; j < input[i].length; j++) {
    const current = input[i][j];

    if (current === "X") {
      // check right
      if (input[i][j + 1] === "M" && input[i][j + 2] === "A" && input[i][j + 3] === "S") {
        count++;
      }

      // check left
      if (input[i][j - 1] === "M" && input[i][j - 2] === "A" && input[i][j - 3] === "S") {
        count++;
      }

      // check down
      if (input[i + 1]?.[j] === "M" && input[i + 2]?.[j] === "A" && input[i + 3]?.[j] === "S") {
        count++;
      }

      // check up
      if (input[i - 1]?.[j] === "M" && input[i - 2]?.[j] === "A" && input[i - 3]?.[j] === "S") {
        count++;
      }

      // check diagonal right down
      if (input[i + 1]?.[j + 1] === "M" && input[i + 2]?.[j + 2] === "A" && input[i + 3]?.[j + 3] === "S") {
        count++;
      }

      // check diagonal right up
      if (input[i - 1]?.[j + 1] === "M" && input[i - 2]?.[j + 2] === "A" && input[i - 3]?.[j + 3] === "S") {
        count++;
      }

      // check diagonal left up
      if (input[i - 1]?.[j - 1] === "M" && input[i - 2]?.[j - 2] === "A" && input[i - 3]?.[j - 3] === "S") {
        count++;
      }

      // check diagonal left down
      if (input[i + 1]?.[j - 1] === "M" && input[i + 2]?.[j - 2] === "A" && input[i + 3]?.[j - 3] === "S") {
        count++;
      }
    }
  }
}

console.log(count);
