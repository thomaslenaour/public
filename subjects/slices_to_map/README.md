## slices_to_map

### Instructions:

Create a function that borrows two slices and returns a hashmap where the first slice represents the keys and the second represents the values.

### Expected Function

```rust
pub fn slices_to_map(&[T], &[U]) -> HashMap<&T, &U> {
}
```

### Usage

Here is a program to test your function.

```rust
fn main() {
	let keys = ["Olivia", "Liam", "Emma", "Noah", "James"];
	let values = [1, 3, 23, 5, 2];
	println!("{:?}", slices_to_map(&keys, &values));
}
```

And its output

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
{"Liam": 3, "James": 2, "Emma": 23, "Noah": 5, "Olivia": 1}
student@ubuntu:~/[[ROOT]]/test$
```