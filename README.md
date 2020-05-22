# HRM Compiler
_Be sure to check out the game on their [website](https://tomorrowcorporation.com/humanresourcemachine) or on [Steam](https://store.steampowered.com/app/375820/Human_Resource_Machine/)._

## Usage
`hrm <level> <source path>`
- Level is the in-game level number you want to test for
- Source path is the location of the code copied from/to be pasted into the game
`comments --decode <path | text>`
- Path to base64-encoded HRM comment file, or as text
- Generates an image `out.png` visualizing the comment
`comments --encode <text>`
- Encodes up to 26 characters of UPPERCASE characters to generate a comment
- Writes encoded form to stdout

## Features
- Complete compiler for the Human Resource Machine (HRM) language
- Debugging tools for developing the compiler
- Levels 1-22* with deterministic testing (differs from in-game tests, but covers all possible edge cases)
- Encoding and decoding of the comment system using drawings
