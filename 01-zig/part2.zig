const std = @import("std");
const input = @embedFile("data");

pub fn main() !void {
    // Iterator over the string split across new lines.
    var iterator = std.mem.tokenize(u8, input, "\n");

    // Initializing the counter.
    var counter: usize = 0;

    // Create the initial window.
    var window: [3]usize = .{
        std.fmt.parseInt(usize, iterator.next().?, 10) catch return,
        std.fmt.parseInt(usize, iterator.next().?, 10) catch return,
        std.fmt.parseInt(usize, iterator.next().?, 10) catch return,
    };

    // Variable to hold the sum.
    var old_sum: usize = sumArr(window);

    while (iterator.next()) |val| {
        // Move the window down by 1.
        window[0] = window[1];
        window[1] = window[2];
        window[2] = std.fmt.parseInt(usize, val, 10) catch return;

        var new_sum: usize = sumArr(window);

        // Check the sum
        if (new_sum > old_sum) {
            counter += 1;
        }

        // Swap the sum values.
        old_sum = new_sum;
    }
    std.debug.print("Result: {}\n", .{counter});
}

// Just a simple function to add all elements of an array into 1.
fn sumArr(arr: [3]usize) usize {
    var sum: usize = 0;
    for (arr) |val| {
        sum += val;
    }
    return sum;
}
