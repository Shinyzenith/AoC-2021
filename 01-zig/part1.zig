const std = @import("std");
const input = @embedFile("data"); // embed the entire freaking file in the execuatable.

pub fn main() !void {
    // Iterator over the string split across new lines.
    var iterator = std.mem.tokenize(u8, input, "\n");

    // Pop the first value by calling iterator.next() and parse to int.
    var last: usize = std.fmt.parseInt(usize, iterator.next().?, 10) catch return;

    // Counter to keep a track of the result.
    var counter: usize = 0;

    while (iterator.next()) |line| {
        // Try to parse string to int.
        const n = std.fmt.parseInt(usize, line, 10) catch continue;

        // increment counter.
        if (n > last) counter += 1;

        // Swap.
        last = n;
    }
    std.debug.print("Result: {}\n", .{counter});
}
