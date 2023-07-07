# Website Status Monitor

This project was based on the Alura course: https://cursos.alura.com.br/course/golang.

I made some changes compared to the design of the course to make it more in line with my style, this was my first contact with Golang so there are certainly changes that can be made to make it more enjoyable.

The unit tests were done exclusively by me, researching how to do unit tests in golang, looking for the official documentation. Finally, I discovered that it is not difficult to create tests from scratch in Go, being possible to carry out tests without having to download any extra dependencies or frameworks.

There are definitely gaps where I can improve the tests and apply better logic, including in the executable script. I can see that, because I had to make changes to all the functions so that it was possible to test their return, with that the tests worked perfectly, they were even tested applying mutation testing techniques to make sure they were well implemented.

To run the program, you must download the golang dependencies and run the command: go run check_site_status.go
To run the tests, run the command: go test . -cover

Running this way with ".", it will only return the log of tests that fail and with the flag" -cover" it will be possible to see the total coverage of the test
