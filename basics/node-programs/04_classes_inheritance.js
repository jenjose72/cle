/*
	NODE.JS PROGRAM 4: CLASSES & INHERITANCE
	Concepts:
	- Using 'class' for object-oriented design.
	- Constructor to initialize values.
	- Methods to add behavior.
	- Basic inheritance ('extends').
*/

// Shape Class (Base)
class Shape {
	constructor(name) {
		this.name = name;
	}
	
	// Method that can be overridden
	area() {
		return 0;
	}
}

// Rectangle Class (Inherits Shape)
class Rectangle extends Shape {
	constructor(width, height) {
		super("Rectangle"); // Call the parent constructor
		this.width = width;
		this.height = height;
	}
	
	area() {
		return this.width * this.height;
	}
}

// Circle Class (Inherits Shape)
class Circle extends Shape {
	constructor(radius) {
		super("Circle");
		this.radius = radius;
	}
	
	area() {
		return 3.14 * this.radius * this.radius;
	}
}

// Logic
const r = new Rectangle(10, 5);
const c = new Circle(7);

console.log("Working with Shapes (Classes):");
console.log(`${r.name} Area: ${r.area()}`);
console.log(`${c.name} Area: ${c.area()}`);
