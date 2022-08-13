# Taxi Driver Matcher
 
A Golang implementation of REST API service for matching taxi drivers with pedestrians. QuadTree data structure for matching was used. The packages are covered with tests by **99%**

-------------------

### Driver searching benchmark
The positions of the driver and pedestrian are randomly distributed on the map.

```
BenchmarkSearch/100_drivers_100000_mapdimension-8         	  861868	      1329 ns/op	      72 B/op	       3 allocs/op
BenchmarkSearch/1000_drivers_100000_mapdimension-8        	  767870	      1554 ns/op	      75 B/op	       3 allocs/op
BenchmarkSearch/10_drivers_1000000_mapdimension-8         	 1000000	      1247 ns/op	      61 B/op	       2 allocs/op
BenchmarkSearch/1000000_drivers_1000_mapdimension-8       	  672942	      1694 ns/op	     156 B/op	       6 allocs/op
```