# fav-imgs
A small sample project of a web system in Golang that stores and displays "my favorite pictures".
My main objective is to demonstrate a clean code and clean architecture approach. 🧱

#### Notice how easy it was to implement the last features (delete and modify image), without the risk of breaking anything. The structure of the code is SOLID.

The next step would be to generalize the front-end by making it a Json (probably graphQL) interface (and build a new, independently deployable front-end)... and to make the persistence implementation into a plugin, but for now plugins are not yet available in windows for Go (Unix only), so I will wait for them before I do it.

## how to install
Just download and build with Golang using `go build`, then run the executable. It will start a server on `http://localhost:8080`
- To see the main page (with the list of images) go to `http://localhost:8080`. You should be able to access all functionality from here.

### Important note
This version is only for local deployment and testing because of some hard-coded values. I will update it to make it web-deployable in the future.

## Implemented
- List all images
- Add image
- Delete image
- Modify image


## File Structure
You will find the main method on `main.go` and the main server handler functions on `/server`.

⚠️**The core code is inside the `/gallery` folder.** 

## Architectural notes
- TDD, 100% coverage for the back-end
- Instead of using an existing Database technology, I created an AdHoc persistance using Json and a text file (a simple 
database), which is a lot easier to deploy (no need to set up SQL database), and works well with the current 
requirements. For this I use a map with automatically generated strings for IDs.
- We can further separate the front-end from the back-end by changing the back-end to instead return Json structures as 
an API, and then building the front-end separately (using React, for example, to get a more dynamic page). That will 
also make them independently deployable, which will in turn make it a lot easier to build an Android or iOS client App, 
for example. I did not do this because I did not want to go back and forth between two different stacks at the beginning 
of the development. I will consider doing this in a future iteration.
- We follow the Stable Dependencies Principle and Stable Abstractions Principle for coupling.
- You can modify the front-end templates without needing to redeploy the software.
- I can add the option to add either an image URL or an image file in the future, but for now I wanted to focus on 
demonstrating the architecture, so you can only add images by url.
- The front-end is relatively messy because it is the least stable component.
- I did not add CSS because it is irrelevant for the purpose of this exercise.  🥱

### Important note
Some refactor is still needed, but the code is kept clean.

#### There are not a lot of comments because the code does not need a lot of comments.
