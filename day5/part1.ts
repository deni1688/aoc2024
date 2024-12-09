import * as fs from 'fs';

let [rawRules, rawReports] = fs.readFileSync('day5/input.txt', 'utf-8').split('\n\n')

const rules= rawRules.split('\n').map(y => y.split('|').map(z => Number(z)))
const reports = rawReports.split('\n').map(x => x.split(',').map(y => Number(y)))

console.log({rules});
console.log({reports});
