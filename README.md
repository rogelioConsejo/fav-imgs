# fav-imgs
A small sample project of a web system that displays "my favorite pictures"

## how to install
Just download and build with Golang using `go build`, then run the executable. It will start a server on `http://localhost:8080`
- To add an image go to `http://localhost:8080/add`
- To list existing images goto `http://localhost:8080/list`

## Implemented
- List all images
- Add image

## Pending Implementation
- Delete image
- Modify image

## File Structure
You will find the main method on `main.go` and the main server handler functions on `/server`.

The core code is inside the `/gallery` folder.

## Architectural notes
- TDD, 100% coverage for the back-end
- Instead of using an existing Database technology, I created an AdHoc persistance using Json and a text file (a simple database), which is a lot easier to deploy (no need to set up SQL database), and works well with the current requirements. For this I use a map with automatically generated strings for IDs.
- We can further separate the front-end from the back-end by changing the back-end to instead return Json structures as an API, and then building the front-end separately (using React, for example, to get a more dynamic page). That will also make them independently deployable, which will in turn make it a lot easier to build an Android or iOS client App, for example. I did not do this because I did not want to go back and forth between two different stacks at the beginning of the development. I will consider doing this in a future iteration, after implementing all 4 basic behaviours.
- We follow the Stable Dependencies Principle and Stable Abstractions Principle for coupling
