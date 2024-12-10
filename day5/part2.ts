import * as fs from 'fs';
import * as process from "node:process";

let [rawRules, rawUpdates] = fs.readFileSync('day5/input.txt', 'utf-8').trim().split('\n\n')

const rules= rawRules.split('\n').map(y => y.split('|').map(z => Number(z)))
const updates = rawUpdates.split('\n').map(x => x.split(',').map(y => Number(y)))

function isBad(row: number[]) {
    for (let [left, right] of rules) {
        let indexOfLeft = row.indexOf(left);
        let indexOfRight = row.indexOf(right);

        if (indexOfLeft !== -1 && indexOfRight !== -1 && indexOfLeft > indexOfRight) {
            row[indexOfLeft] = right;
            row[indexOfRight] = left;

            return true;
        }
    }

    return false;
}

function processRow(badRowSet: Set<number>,row: number[], index: number) {
    if (isBad(row)) {
        badRowSet.add(index);
        processRow(badRowSet, row, index);
    }

    return badRowSet;
}

const count = Array.from(updates.reduce(processRow, new Set<number>())).reduce((acc, rowIndex) => {
    const row = updates[rowIndex];
    acc += row[ Math.floor(row.length / 2)];

    return acc;
}, 0);


console.log(count);


