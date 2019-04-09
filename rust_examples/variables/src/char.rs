fn main() {
   let c = 'z';
   println!("c is : {}", c);
   let s = String::from("hello");
   println!("String is : {}", s);
   let mut s1 = String::from("hello");
   s1.push_str(", world!"); // push_str() appends a literal to a String
   println!("{}", s1); // this will print `hello, world!`
}