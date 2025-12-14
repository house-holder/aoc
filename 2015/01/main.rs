use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        eprintln!("Usage: {} <filename>", args[0]);
        return;
    }

    let filename = &args[1];
    let input = fs::read_to_string(filename).expect("Failed to read file");

    let mut floor = 0;
    for ch in input.chars() {
        match ch {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => {}
        }
    }
    println!("Part 1: {}", floor);

    let mut floor = 0;
    for (i, ch) in input.chars().enumerate() {
        match ch {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => {}
        }
        if floor == -1 {
            println!("Part 2: {}", i + 1);
            break;
        }
    }
}
