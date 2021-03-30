export const decodedValue = (colors) => {
  const firstTwo = colors.slice(0, 2)
  return (10 * COLORS.indexOf(firstTwo[0])) + COLORS.indexOf(firstTwo[1])
} 

export const COLORS = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white',
]
