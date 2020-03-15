from __future__ import print_function
from ortools.constraint_solver import pywrapcp
from os import abort

def CPIsFun():
  # Constraint programming engine
  solver = pywrapcp.Solver('CP is fun!');
  
  kBase = 10
  
  # Decision variables.
  digits = list(range(0, kBase))
  digits_without_zero = list(range(1, kBase))
  a = solver.IntVar(digits_without_zero, 'A');
  b = solver.IntVar(digits, 'B');
  c = solver.IntVar(digits_without_zero, 'C');
  d = solver.IntVar(digits, 'D');
  e = solver.IntVar(digits, 'E');
  f = solver.IntVar(digits, 'F');
  g = solver.IntVar(digits, 'G');
  h = solver.IntVar(digits, 'H');
  i = solver.IntVar(digits_without_zero, 'I');
  j = solver.IntVar(digits_without_zero, 'J');

  # We need to group variables in a list to use the constraint AllDifferent.
  letters = [a,b,c,d,e,f,g,h,i,j]
  diffletters = [a, b]
  placeholders = [c,d,e,f,g,h,i,j]
  
  # Verify that we have enough digits.
  assert kBase >= len(diffletters)

  # Define constraints.
  solver.Add(solver.AllDifferent(diffletters))
    
  # AGHB + IAB0 = JBAB
  solver.Add (b + kBase * (h+b) + kBase * kBase * (g+a) + kBase * kBase * kBase * (a+i) ==
              b + kBase * a + kBase * kBase * b + kBase * kBase * kBase * j) 
  # CDE * AF = JBAB
  solver.Add ( (e + kBase * d + kBase * kBase * c) * (f + kBase * a) ==  b + kBase * a + kBase * kBase * b + kBase * kBase * kBase * j )
  # CDE * F = AGHB
  solver.Add ( (e + kBase * d + kBase * kBase * c) * f == b + kBase * h + kBase * kBase * g + kBase * kBase * kBase * a)
  # CDE * A = IAB
  solver.Add ( (e + kBase * d + kBase * kBase * c) * a == b + kBase * a + kBase * kBase * i)

  db = solver.Phase(letters, solver.INT_VAR_DEFAULT,
                             solver.INT_VALUE_DEFAULT)
  solver.NewSearch(db)
  
  while solver.NextSolution():
    print(letters)
    # Is CP + IS + FUN = TRUE?
    #assert (kBase*c.Value() +  p.Value() + kBase*i.Value() + s.Value() +
            #kBase*kBase*f.Value() + kBase*u.Value() + n.Value() == 
            #kBase*kBase*kBase*t.Value() + kBase*kBase*r.Value() + 
            #kBase*u.Value() + e.Value()) 
  
  solver.EndSearch()

  return


if __name__ == '__main__':
  CPIsFun()
