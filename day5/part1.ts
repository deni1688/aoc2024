import * as fs from 'fs';

let [rawRules, rawUpdates] = fs.readFileSync('day5/input_sample.txt', 'utf-8').trim().split('\n\n')

const rules= rawRules.split('\n').map(y => y.split('|').map(z => Number(z)))
const updates = rawUpdates.split('\n').map(x => x.split(',').map(y => Number(y)))

let middlePageCount = 0;

for (let row of updates) {
    let valid = true;

    for (let [left, right] of rules) {
        let indexOfLeft = row.indexOf(left);
        let indexOfRight = row.indexOf(right);

        if(indexOfLeft !== -1 && indexOfRight !== -1 && indexOfLeft > indexOfRight) {
            valid = false;
            break;
        }
    }

    if(valid) {
        let middleIndex = Math.floor(row.length / 2);
        middlePageCount += row[middleIndex];
    }
}

console.log(middlePageCount);


