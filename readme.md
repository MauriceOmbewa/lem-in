# LEM-IN
## Overview

The goal of this project is to create a digital version of an ant farm simulation in Go, where ants find the quickest path across a network of interconnected rooms. You will implement a program called lem-in that reads a colony structure from a file, identifies the shortest paths from a start room to an end room, and outputs each ant’s movement step-by-step.
## Objectives

    Simulate an Ant Colony: Develop an ant farm with rooms and tunnels, where ants navigate from a start room to an end room.
    Find the Quickest Path: Calculate the optimal path(s) to transport n ants to the end room with minimal moves.
    Handle Complex Colony Configurations: Process various room and tunnel arrangements, some of which may have no viable path between start and end.
    Error Handling: Account for invalid inputs (e.g., incorrect room formatting, missing start/end, loops, etc.) and display appropriate error messages.

## Program Requirements

    Input Structure:
        Input file should describe rooms, links, and ants in the following format:

    number_of_ants
    the_rooms
    the_links

    Each room is defined as: name coord_x coord_y
    Each link is defined as: room1-room2

## Output Structure:

    Display file contents and each ant’s movement between rooms in a format such as:

        Lx-y Lz-w Lr-o ...

        Here, x, z, r are ant numbers, and y, w, o are room names.

    Game Rules:
        Ants start in the room labeled ##start and must reach ##end.
        Each ant can only move once per turn and can only occupy an empty room.
        Tunnels are one-way only and cannot be reused within the same turn.

    Constraints:
        Only standard Go packages are allowed.
        Program must be written with good Go practices and should include unit tests.

## Usage

To run the program, execute:

$ go run . [input_file]

Example

Input file:

3
##start
1 23 3
2 16 7
##end
0 9 5
0-4
1-2

Output:

L1-2 L2-3
L1-4 L2-0
L3-1

Error Handling

    Invalid file format or data (e.g., missing rooms, duplicate tunnels) should display an error message:

    ERROR: invalid data format

Project Structure

    Ants and Rooms: Store ants, rooms, and tunnels in appropriate Go structures for efficient pathfinding.
    Algorithm: Implement shortest-path algorithms to ensure ants find the optimal path while avoiding traffic jams.
    Error Handling: Detect and respond to incorrect configurations gracefully.

Bonus
Visualizer

As an optional feature, create an ant farm visualizer that displays ants moving through the colony in real-time:

$ ./lem-in ant-farm.txt | ./visualizer

Learning Outcomes

This project emphasizes:

    Algorithm development for pathfinding.
    Input parsing and structured data handling.
    Efficient data manipulation and output formatting.
    Error handling in Go.

## License

This project is open-source and licensed under the MIT License.

Enjoy coding your digital ant farm and optimizing their paths!