export const twoFer = (name) => {
  const isString = typeof name === 'string' && Object.prototype.toString.call(name) === '[object String]'
  if (!isString || !name) {
    return 'One for you, one for me.'
  }

  return `One for ${name}, one for me.`
};
