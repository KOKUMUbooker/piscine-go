// Empty file
package piscine

func FirstRune(s string) rune {
  var rVal rune
  for i,r := range s {
    if i == 0 {
      rVal = r
      break
    }
  }

  return rVal
}