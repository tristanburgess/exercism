const GIGASECOND = 10 ** 9

export const gigasecond = (date) => {
  return new Date(date.getTime() + GIGASECOND * 1000)
};
