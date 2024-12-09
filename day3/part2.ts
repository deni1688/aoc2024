import * as fs from 'fs';

const regex = /mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)/g;

console.log(fs.readFileSync('day3/input.txt', 'utf8').trim().match(regex).map(m => {
    switch (true) {
        case m === "don't()": {
            return false
        }
        case m === "do()": {
            return true
        }
        default: {
            return eval(m.replace(regex, "$1*$2"))
        }
    }
}).reduce((acc, val) => {
    if (val === false) {
        acc.skip = true;
        return acc;
    }


    if (val === true) {
        acc.skip = false;
        return acc;
    }

    if (acc.skip) return acc;


    acc.total += val;

    return acc;
}, {total: 0, skip: false}).total);
