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
  b = solver.IntVar(digits_without_zero, 'B');
  c = solver.IntVar(digits, 'C');
  d = solver.IntVar(digits_without_zero, 'D');
  e = solver.IntVar(digits, 'E');
  f = solver.IntVar(digits, 'F');
  g = solver.IntVar(digits_without_zero, 'G');
  h = solver.IntVar(digits, 'H');
  i = solver.IntVar(digits, 'I');
  j = solver.IntVar(digits, 'J');
  k = solver.IntVar(digits_without_zero, 'K');
  l = solver.IntVar(digits, 'L');
  m = solver.IntVar(digits, 'M');
  n = solver.IntVar(digits, 'N');
  o = solver.IntVar(digits, 'O');
  p = solver.IntVar(digits, 'P');
  q = solver.IntVar(digits, 'Q');
  r = solver.IntVar(digits_without_zero, 'R');
  s = solver.IntVar(digits, 's');
  t = solver.IntVar(digits, 't');
  u = solver.IntVar(digits, 'u');
  v = solver.IntVar(digits, 'v');
  w = solver.IntVar(digits, 'w');
  x = solver.IntVar(digits, 'x');

  # We need to group variables in a list to use the constraint AllDifferent.
  letters = [a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x]
  
  # Verify that we have enough digits.
  #assert kBase >= len(letters)

  # Define constraints.
  #solver.Add(solver.AllDifferent(diffletters))
  solver.Add(a!=b)
  solver.Add(a!=c)
  solver.Add(a!=d)
  solver.Add(a!=e)
  solver.Add(a!=f)
  solver.Add(a!=g)
  solver.Add(a!=h)
  solver.Add(a!=i)
  solver.Add(a!=j)
  solver.Add(a!=k)
  solver.Add(a!=l)
  solver.Add(a!=m)
  solver.Add(a!=n)
  solver.Add(a!=o)
  solver.Add(a!=p)
  solver.Add(a!=q)
  solver.Add(a!=r)
  solver.Add(a!=s)
  solver.Add(a!=t)
  solver.Add(a!=u)
  solver.Add(a!=v)
  solver.Add(a!=w)
  solver.Add(a!=x)
    
  # BAAC * F = GHAIJ
  solver.Add ( (c + kBase * a + kBase * kBase * a + kBase*kBase*kBase*b) * f ==  j + kBase * i + kBase * kBase * a + kBase * kBase * kBase * h + kBase*kBase*kBase*kBase*g )
  # BAAC * E = KLMNA
  solver.Add ( (c + kBase * a + kBase * kBase * a + kBase*kBase*kBase*b) * e ==  a + kBase * n + kBase * kBase * m + kBase * kBase * kBase * l + kBase*kBase*kBase*kBase*k )
  # BAAC * D = AOPQA
  solver.Add ( (c + kBase * a + kBase * kBase * a + kBase*kBase*kBase*b) * d ==  a + kBase * q + kBase * kBase * p + kBase * kBase * kBase * o + kBase*kBase*kBase*kBase*a )
  # BAAC * DEF = RSTUVWX
  solver.Add ( (c + kBase * a + kBase * kBase * a + kBase*kBase*kBase*b) * (f + kBase*e + kBase*kBase*d) == x +kBase*w +pow(kBase,2)*v +pow(kBase,3)*u + pow(kBase,4)*t +pow(kBase,5)*s + pow(kBase,6)*r)
  # GHAIJ + 10*KLMNA + 100*AOPQA = RSTUVWX
  solver.Add( (j + kBase * i + kBase * kBase * a + kBase * kBase * kBase * h + kBase*kBase*kBase*kBase*g) + kBase*(a + kBase * n + kBase * kBase * m + kBase * kBase * kBase * l + kBase*kBase*kBase*kBase*k) + pow(kBase, 2)*(a + kBase * q + kBase * kBase * p + kBase * kBase * kBase * o + kBase*kBase*kBase*kBase*a) == x +kBase*w +pow(kBase,2)*v +pow(kBase,3)*u + pow(kBase,4)*t +pow(kBase,5)*s + pow(kBase,6)*r)
  
  

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
