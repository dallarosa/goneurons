# goneurons
Basics of Neural Networks in Go

I've decided to start this project as a way to put together enthusiasm for Go
together an opportunity to learn a bit.

This is a very simplistic Neuron. It accepts a 2-dimensional array, n features 
and m examples.

The code should be pretty easy to read in itself.

To try it out, just change initialize the neuron with the input and output of
your liking.

At some point I might move the input and output to be set from the command line
but for now just change it in the main function. the same goes for the number of
iterations, the number of features and the learning rate.

The idea is to keep adding features so that at some point you can actually do some
interesting things, but this is by no means focused on speed or scalability. The 
main focus is code readiability and being a educational resource (mostly for myself).

Things in my TODO list:

* Allowing Neurons to be stacked in layers
* Add Optimization Algorithms (Adams)
* Add Activation Functions (ReLU, Sigmoid, Tanh, Softmax)

I'm also open to suggestions.

DISCLAIMER: This was created for educational purposes. Feel free to use this code
but use it at your own discretion.
