# French Number Converter Go implementation

## Objectives

- Code Clarity and Maintainability: The code prioritizes readability, understandability, and ease of maintenance

  - State machines provide a clear structure for handling different number ranges with specific conversion rules.

- Infrastructure:

  - I provide a docker container to facilitate isolation and portable environment

- Thorough Testing

## Requirements

- Numbers in French have complexities due to various rules depending on their range and digit composition.
- This program addresses the core rules for writing numbers in French, referencing [Nombres en français](https://fr.wikipedia.org/wiki/Nombres_en_fran%C3%A7ais).

**Running the Implementation**

1.  **Build, Test and Run the Application:**

I provided a Makefile to simplify, running, testing and building the application, to this end you can run the implementation

- Run and build

  - Open a terminal

    - `git clone https://github.com/GroundZeroGreyHead/jusmundi.git`

    - `cd jusmundi`

    - To run the docker version app 👉 `make run-docker`
    - To run the on local OS app 👉 `make build-and-run`

- Running tests

        make test

Don't forget to clean up.

    make clean

🍻 Cheers!

Mubarak
