import os

pub fn main() {
	file_path := "./data"

	// Read the file and split by new lines.
	lines := os.read_file(file_path)?.split_into_lines()

	mut horizontal := 0
	mut depth := 0

	for line in lines {
		input := line.split(" ")
		travel_direction := input[0].trim(" ")
		travel_distance := input[1].int()

		match travel_direction {
			"forward" {
				horizontal += travel_distance
			} 
			"down" {
				depth += travel_distance
			}
			"up" {
				depth -= travel_distance
			}
			else {} // It's already exhaustive.
		}
	}
	println("Result: ${horizontal*depth}")
}
