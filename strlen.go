package piscine

func StrLen(s string) int {
  l := 0
  for i,_ := range s {
      l = i
  }
  
  return l + 1
}