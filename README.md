# hello-goapp

There are two services in the project:

* hellofoo
* hellobar

hellofoo service has different directory structure for hybrid app support from hellobar which is simple and has only app.go

To run hellofoo on dev server:

    goapp serve hellofoo/app


To run hellobar on dev server with different port:

    goapp serve -port=9090 -admin_port=9091 hellobar
    

### Trace test

Let's see the trace information on google cloud console.
Call the following url via 'GET' method, so that the call from hellofoo fetchs the hellobar's url.
Then, the information should be shown on the trace console.

    http://localhost:8080/trace-foo
