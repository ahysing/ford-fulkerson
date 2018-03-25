Ford-Fulkerson Algorithm
========================

This program finds the maximum capacity in a directed graph with capacity at every edge. This is known as the Ford-Fulkerson algorithm.

It traverses from a start node named `source` to the final node named `sink` by finding a valid path through the graph. As it finds a valid path it allocates the maximum flow through that path to the entire graph.
According to the [Max-flow min-cut theorem](https://en.wikipedia.org/wiki/Max-flow_min-cut_theorem) any allocation of capacity which has exausted all available flows is a maximum flow.

When all paths are exausted the maximum flow is added up by summing all capacity allocated from the start node.

Tech
----
This program is written, and debugged with Go 1.10
