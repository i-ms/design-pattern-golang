# Bridge Pattern

**Introduction**
Bridge pattern allows us to split a large class or a set of closely related classes into separate hierarchies - abstraction and implementation - which can be developed independently.
One of these hierarchies (often called the abstraction) will get a reference to an object of the second hierarchy (implementation). The abstraction will be able to delegate some (sometimes, most) of its calls to the implementations objects.
