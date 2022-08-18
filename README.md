## Circutor Technical Test

This is a very simple application that shows some energy data Helsinki buildings from 1st January 2021 to around 28 Feb 2022.

It has an API written in GO and a frontend written in React.

## Launch

Clone/Fork the repo with git clone <url>
Build the docker containers and launch the project with:
```sh
docker-compose up --build
```

The api is accessible on `http://localhost:1234/docs/index.html`
The frontend is accessible on `localhost:3001`

## Tasks

- Implement an error message when login fails
- Implement a small test for this new feature

Feel free to implement any other improvement as long as you write a test for it.

## List Redflags

Write here all red flags that you find in the code. Any examples that would stop a code review. If you want to fix some of them, go on.


## How you would make this application maintainable and scalable

Write here all the steps you would take.

## Test submission

Please, submit this test as a new repository (a fork or a new one) in any free platform you want (bitbucket, gitlab, github, ..)