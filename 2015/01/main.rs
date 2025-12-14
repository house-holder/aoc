#![allow(non_snake_case)]

use std::env;
use std::fs;

fn evalPart1(input: &String) -> i32 {
    let mut floor = 0;
    for ch in input.chars() {
        match ch {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => {}
        }
    }
    return floor;
}

fn evalPart2(input: &String) -> i32 {
    let mut floor = 0;
    for (i, ch) in input.chars().enumerate() {
        match ch {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => {}
        }
        if floor == -1 {
            return (i + 1) as i32;
        }
    }
    return -1;
}

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        eprintln!("Usage: {} <filename>", args[0]);
        return;
    }
    let filename = &args[1];
    let input = fs::read_to_string(filename).expect("Failed to read file");

    let result1 = evalPart1(&input);
    println!("Part 1: {}", result1);
    let result2 = evalPart2(&input);
    println!("Part 2: {}", result2);
}
