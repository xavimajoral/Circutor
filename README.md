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
- Wrap all the functions (i.e: /app/src/api/auth.js) that fetch something from the API with: 

```
try {

} catch (error) {
  // Do sth with the error
}
```

- Some fetches from the API are very similar. We could build a custom react hook (i.e: useFetch) for the calls within ```web/src/api```
- login component and register component are almost the same. We could try to make just 1 component for both of them.
- Some functions are typed in different files: logoutUser, goToBookrmarks, goToBuildings, goToBuildingMetrics. Put them in 1 file and import them.
- There are 2 routes with the same component being rendered. Not sure if there's a reason for that
<Route path="home" element={<LayoutRegister />} />
<Route path="/" element={<LayoutRegister />} />
- Some unused imports at the top of some files
- There's a warning: ```Warning: Failed prop type: MUI: You are providing an onClick event listener to a child of a button element.
Prefer applying it to the IconButton directly.
This guarantees that the whole <button> will be responsive to click events.```
But I'd say is wrong and the onClick event is properly added. Some other users complaining about this issue.
- Instead of using Moment library we can build the date easily:
`new Date().toISOString().slice(0,10); //return YYYY-MM-DD` or also use date.getDate(), date.getMonth(), date.getFullYear() and build the string.
- I'd ask what's the purpose of ```navigate(0)``` that appears along the code.
- Large render methods


## How you would make this application maintainable and scalable

- I'd try to avoid code duplication as much as possible to build reusable components (i.e: login/logout, header)
- In terms of how to apply css different techniques are used. I'd try to define a style guide within the team and follow all of us the same practice.
- Setup unit testing for the components covering a pre-defined percentage of covering (i.e: > 75%)
- Setup end to end test. I'd go along this task with Cypress
- Having tools that allows us to have Continuous Integration where we can check: code coverage, page speed, clean code
- When needed add comments why you do that in particular.
- Option to add Typescript so that we can find some errors on pre-compilation time.

## Test submission

Please, submit this test as a new repository (a fork or a new one) in any free platform you want (bitbucket, gitlab, github, ..)