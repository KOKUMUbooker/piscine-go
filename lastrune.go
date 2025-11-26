package piscine

func LastRune(s string) rune {
  var lastChar rune
  for i,r := range s {
    if i == len(s) - 1 {
      lastChar = r
    }
  }
  return lastChar
}