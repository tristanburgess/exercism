const complements = {
  "G": "C",
  "C": "G",
  "T": "A",
  "A": "U",
}

// Loop impl
/*
export const toRna = (dnaSequence) => {
  let rnaSequence = ""
  for (let n of dnaSequence) {
    rnaSequence += complements[n]
  }
  return rnaSequence
};
*/

// Reduce impl
/*
export const toRna = (dnaSequence) => {
  return [...dnaSequence].reduce((rnaSequence, n) => rnaSequence += complements[n], '')
};
*/

// Map/Join impl
export const toRna = (dnaSequence) => {
  return [...dnaSequence].map((n) => complements[n]).join('')
};
