function* generVal(val, multiplyVal, multiple) {
  while (true) {
    val = (val * multiplyVal) % (2 ** 31 - 1);
    if (val % multiple == 0) {
      yield val;
    }
  }
}

let pair = 0;

let aGen = generVal(873, 16807, 4);
let bGen = generVal(583, 48271, 8);

for (let i = 0; i < 5000000; i++) {
  aVal = aGen.next().value;
  bVal = bGen.next().value;

  a16 = aVal & 0xffff;
  b16 = bVal & 0xffff;

  pair += a16 == b16;
}

console.log(pair);
