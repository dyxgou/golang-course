# Map

The golang maps implementation is created by a hash map.

A hash map is a kind of data structure that is made with an array and a hash function, there are a lot of hash functions, but their main porpuse is to take de value of that is about to be put in tthe map, and generate a random deterministric key.


```golang
func HashFunc(value interface{}) Key
```
