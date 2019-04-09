fn main() {

    let guess: u32 = "40".parse().expect("Not a number!");
    // let guess: u32 = "42a".parse().expect("Not a number!");
    println!("There are {} spaces", guess);
}