class Squares(val n: Int) {
  fun squareOfSum(): Int {
    val sum = (1..n).asIterable().sum()
    return sum * sum
  }

  fun sumOfSquares(): Int {
    return (1..n).asIterable().reduce { s, num -> s + num * num }
  }

  fun difference(): Int {
    return squareOfSum() - sumOfSquares()
  }
}