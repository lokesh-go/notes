# Design Patterns

- Design patterns are typical solutions to commonly occurring problems in software design.
- They are like pre-made blueprints that you can customize to solve a recurring design problem in your code.

---

#### The Catalog of Design Patterns

##### Creational patterns

These patterns provide various object creation mechanisms, which increase flexibility and reuse of existing code.

- Factory Method
  - Factory Method is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.
  - Factory method is a creational design pattern which solves the problem of creating product objects without specifying their concrete classes.
- Abstract Factory
  - Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying their concrete classes.
  - Abstract Factory is a creational design pattern, which solves the problem of creating entire product families without specifying their concrete classes.
  - Abstract Factory defines an interface for creating all distinct products but leaves the actual product creation to concrete factory classes. Each factory type corresponds to a certain product variety.
- Builder
  - Builder is a creational design pattern that lets you construct complex objects step by step.
  - The pattern allows you to produce different types and representations of an object using the same construction code.
  - Builder is a creational design pattern, which allows constructing complex objects step by step.
- Singleton
  - Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.
  - Singleton is a creational design pattern, which ensures that only one object of its kind exists and provides a single point of access to it for any other code.
- Prototype
  - Prototype is a creational design pattern that lets you copy existing objects without making your code dependent on their classes.

---

##### Structural Design Patterns

Structural design patterns explain how to assemble objects and classes into larger structures, while keeping these structures flexible and efficient.

- Adapter
  - Also known as: Wrapper
  - Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate.
- Facade
  - Facade is a structural design pattern that provides a simplified interface to a library, a framework, or any other complex set of classes.
- Decorator
  - Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors.


---

##### Behavioral Design Patterns

Behavioral design patterns are concerned with algorithms and the assignment of responsibilities between objects.

- Observer
  - Observer is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about any events that happen to the object they’re observing.
- Strategy
  - Strategy is a behavioral design pattern that lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.
  - Strategy is a behavioral design pattern that turns a set of behaviors into objects and makes them interchangeable inside original context object.