# Leetcode Helper

This is a simple CLI tool that helps you generate test cases for Leetcode questions.

## Note
This is a personal project and is not affiliated with Leetcode.
This uses https://alfa-leetcode-api.onrender.com to get the daily question.

## Usage

1. Run `go build` to build the binary or download the binary from the [release page](https://github.com/Ar0manKhan/leetcode-helper/releases).
2. Get your Gemini API key from [google ai studio](https://aistudio.google.com/) and set it as an environment variable `GEMINI_API_KEY`.
3. Install fmt library for c++ using [this instruction](https://fmt.dev/latest/get-started/).
4. (Optional) Save the binary to a directory in your path.
5. Run `./leetcode-helper` to get the question and generate the test case in the directory where you want to save the test code.

## Features
- Get the daily question from Leetcode in C++ and save it to a file.

## TODO
- [ ] Add support for other languages.
- [ ] Add support for other types of questions.
