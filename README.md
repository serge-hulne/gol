The aim of this experiment is to attempt to model a "parallel" (concurrent) cellular automaton (as opposed to a sequential "generation-based" cellular automaton):



Conventional implementations of cellular automata, in particular implementations of Conway's game-of-life, seem to proceed from the idea of generations (or cycles), wherein the state of the cells making up the automaton remain "frozen" during each "generation" or "cycle" or "iteration", and are changed, in one go,  every cycle, in a loop.

I do not know if it makes any sense from a mathematical or logical point of view, but it certainly looks odd when compared to a (even ultra-simplified) real-world heap of cells:



Let's imagine for a moment that one has (theoretical) living cells, say e.g. genetically engineered precisely to follow the (one variation or another) of the game of life !

- Said cells would of course check for the state of their closest neighbors and modify their state accordingly, in an endless cycle.

- However, they most definitely would not wait for the others (i.e. stay "frozen" for an entire generation), nor would they operate at a pre-determined, externally forced pace (i.e. the duration of a given predetermined "cycle").




So the "conventional" (or "sequential" or "generation-based") approach to the game of life just seems inadequate (if not erroneous) to me.  


In other words the conventional approach will almost certainly fail to generate some of the states of the array of cells which can occur when the cell manage themselves independently in their own goroutines.


This is therefore an attempt to implement a (simple) concurrent, generation-less, game-of-life automaton using the goroutines and channels build-in primitives from the Go programming language (which seems easier than to code such an application using conventional POSIX threads).

Serge Hulne.  


