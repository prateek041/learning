## Mutable vs Immutable Objects
- **Mutable Objects**: Can be changed in place. If you modify them, the changes are reflected to the original object.
	- Example:
		- Lists
		- Dictionaries
		- Sets
		- Custom Objects
- **Immutable Objects**: Cannot be changed in place. If you try to modify them, python will create a new object and the original remains unchanged.
	- Examples:
		- Integers
		- Floats
		- Strings
		- Tuples
		- Booleans
		- Frozensets

## Mutable Default Argument Pitfall
When you define a function with default argument that's mutable, python creates the default argument once during function definition and then any modification to that argument persists across function calls.

