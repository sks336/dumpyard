fn main() {
	let tup:(u32, i64, f64, char) = (128, -127, 345.789, 'D');
    println!("(x, y, z, d) = {:?}", tup);
    let (_x,_y,_z,_d) = tup;
    println!("value for char is : {}", _d);
    println!("value for char is : {}", tup.3);
}