const complements = {
  "G": "C",
  "C": "G",
  "T": "A",
  "A": "U",
}

export const toRna = (dnaSequence) => {
  let rnaSequence = ""
  for (let n of dnaSequence) {
    rnaSequence += complements[n]
  }
  return rnaSequence
};
