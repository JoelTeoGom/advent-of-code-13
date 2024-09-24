# Advent of Code 2023 - Day 13: Point of Incidence

Welcome to my repository for **Advent of Code 2023**, specifically for **Day 13: Point of Incidence**.

## What is Advent of Code?

[Advent of Code](https://adventofcode.com) is an annual coding challenge that takes place every December. Each day, a new puzzle is released, and coders from all over the world try to solve it. It's a great way to improve coding skills, explore new algorithms, and learn new programming languages.

## About Day 13: Point of Incidence

In this puzzle, we are tasked with navigating through a valley filled with mirrors, analyzing patterns of ash and rocks to find reflection lines. The goal is to locate either a **vertical** or **horizontal line** of symmetry within each pattern and calculate a result based on the positions of these lines. Each reflection helps us understand more about the mirror structure, which we need to analyze carefully.

### Puzzle Description:

- We are given a pattern of ash (`.`) and rocks (`#`) in a grid.
- The task is to identify **reflection lines** within these grids—either vertical (between columns) or horizontal (between rows).
- Based on the reflection lines, we summarize the patterns by adding the number of columns to the left of a vertical reflection line or the number of rows above a horizontal reflection line.

For example, here’s a sample grid:
** #.##..##. ..#.##.#. ##......# ##......# ..#.##.#. ..##..##. #.#.##.#. **

In this case, the vertical reflection line is found between the 5th and 6th columns. We summarize the reflection points and calculate the result accordingly.

### Link to the Puzzle:

You can find the full puzzle description and challenge details here:  
[Advent of Code 2023 - Day 13](https://adventofcode.com/2023/day/13)
also the input: (https://adventofcode.com/2023/day/13/input)
## Why Golang?

I'm using **Golang (Go)** to solve this year's puzzles because it's an efficient, statically-typed language that's great for performance-critical applications. Go's concurrency features and simplicity make it a fun challenge for these kinds of problems, especially when optimizing algorithms or working with performance-based tasks.

### What I'm Learning:

- **Golang syntax** and core features like slices, maps, and string manipulation.
