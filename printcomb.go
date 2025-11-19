// Empty file
package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
  var a,b,c rune;

  for i := '0'; i <= '9'; i++ {
    a = i;

    for j := '0'; j <= '9'; j++ {
      b = j;

      for k := '0'; k <= '9'; k++ {
        c = k;

        if a == b && b == c {
          continue;
        }

        if a < b && b < c {
          z01.PrintRune(a);
          z01.PrintRune(b);
          z01.PrintRune(c);
          z01.PrintRune(',');
          z01.PrintRune(' ');
        }
      }
    }
  }
}