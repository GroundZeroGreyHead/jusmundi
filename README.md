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

Note to help with testing:

      - Used chatgpt to generate a map[int]string that was used in numToFrench

      prompt:

      starting from 20 to 70 return a static go map[int]string key is num from 20 to 70 and value is the corresponding french translation don't skip any values

      => numToFrench := map[int]string{
        20: "vingt",
        21: "vingt-et-un",
        22: "vingt-deux",
        23: "vingt-trois",
        24: "vingt-quatre",
        25: "vingt-cinq",
        26: "vingt-six",
        27: "vingt-sept",
        .
        .
        .
      }

Don't forget to clean up.

    make clean

🍻 Cheers!

Mubarak
