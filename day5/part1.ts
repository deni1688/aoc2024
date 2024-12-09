import * as fs from 'fs';

let [rawRules, rawUpdates] = fs.readFileSync('day5/input.txt', 'utf-8').split('\n\n')

const rules= rawRules.split('\n').map(y => y.split('|').map(z => Number(z)))
const updates = rawUpdates.split('\n').map(x => x.split(',').map(y => Number(y)))

console.log({rules});
console.log({updates});


