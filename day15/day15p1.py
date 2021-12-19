#!/usr/bin/env python3

ip = "input_test.txt"
ip = "input.txt"

mp = []
with open(ip) as f:
 for l in f:
  mp += [[int(x) for x in l[:-1]]]
l=len(mp)-1
def mi(n):
 return mp[n[0]][n[1]]
def nei(n):
 na=[]
 for d in [-1, 1]:
  for c in [0, 1]:
   b = list(n)
   b[c] += d
   if b[0]<0 or b[1]<0 or b[0]>l or b[1]>l:
    continue
   na += [tuple(b)]
 return na
def ucs(s=(0, 0), g=(l,l)):
 n = s; pq = {n: mi(n)}; e = set()
 def pop():
  k= sorted([(v, k) for k, v in pq.items()])[0][1]
  return k, pq.pop(k)
 while True:
  n,nv = pop()
  if n == g:
   return nv-mi(s)
  e.add(n)
  for b in nei(n):
   bv=mi(b)+nv
   if b not in pq and b not in e:
    pq[b]=bv
   elif b in pq and pq[b] > bv:
    pq[b]=bv
print(ucs())
