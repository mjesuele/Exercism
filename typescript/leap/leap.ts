export default function isLeapYear(n: number) {
  const isDivisibleBy4 = n % 4 === 0
  const isDivisibleBy100 = n % 100 === 0
  const isDivisibleBy400 = n % 400 === 0

  return isDivisibleBy4 && (!isDivisibleBy100 || isDivisibleBy400)
}
