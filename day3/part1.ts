import * as fs from 'fs';

const regex = /mul\((\d{1,3}),(\d{1,3})\)/g;

console.log(fs.readFileSync('day3/input.txt', 'utf8').trim().match(regex).map(m => {
    return eval(m.replace(regex, "$1*$2"))
}).reduce((acc, val) => {
    acc += val;
    return acc;
}, 0));
