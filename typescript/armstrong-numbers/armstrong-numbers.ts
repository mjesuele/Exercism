export default {
  isArmstrongNumber: (n: number) => {
    const asString = String(n)
    return (
      n ===
      [...asString].reduce(
        (sum, digit) => (sum += Number(digit) ** asString.length),
        0
      )
    )
  }
}
