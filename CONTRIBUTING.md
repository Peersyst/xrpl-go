# Contributing to xrpl-go

## How to Contribute

You can contribute to the project by:

- Reporting bugs
- Suggesting enhancements
- Implementing features
- Writing documentation
- Writing tests

### Reporting Bugs

1. Check existing issues to ensure the bug hasn't already been reported.
2. Use the bug report template when opening a new issue.
3. Include:
   - A clear, descriptive title
   - Steps to reproduce the issue
   - Expected vs. actual behavior
   - Environment details (OS, version, etc.)
   - Any relevant screenshots or error logs

### Suggesting Enhancements

1. Check existing enhancement issues to avoid duplicates.
2. Use the enhancement proposal template.
3. Provide:
   - Detailed description of the proposed feature
   - Potential implementation approaches
   - Potential benefits and use cases

### Making Pull Requests

#### Setup

1. Fork the repository
2. Clone your fork locally
3. Create a new branch for your contribution
4. Make your changes and commit them
5. Push your changes to your fork
6. Open a pull request from your fork to the main repository

#### Development Guidelines

1. Follow the project's coding style and conventions
2. Write clear, concise commit messages
3. Include tests for new features or bug fixes
4. Update documentation as needed
5. Ensure all tests pass before submitting

#### Pull Request Process

1. Push your changes to your fork
2. Open a pull request with:
   - Clear title describing the change
   - Detailed description of modifications
   - Reference any related issues
3. Await code review from maintainers

### Development Setup

#### Prerequisites

- Go 1.22.0 or later
- Go toolchain 1.22.5 or later

#### Installation

```bash
# Clone the repository
git clone https://github.com/Peersyst/xrpl-go

# Install dependencies
go mod tidy

# Run linting
make lint

# Run tests
make test-ci
```

## Style Guide

- Use consistent indentation
- Write clear, self-documenting code
- Add comments for complex logic

## Code Review Process

- All submissions require review from project maintainers
- We may request changes or provide feedback
- Expect a response within [X] business days
- Maintain a respectful and constructive dialogue

## Licensing

By contributing, you agree that your contributions will be licensed under the project's [LICENSE TYPE].

## Questions?

If you have any questions, please:
- Check the documentation
- Open an issue for clarification

**Happy Contributing!** 🚀