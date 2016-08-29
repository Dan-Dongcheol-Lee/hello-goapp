# hello-goapp

There are two apps in the project:

* hellofoo
* hellobar

To run hellofoo on dev server

    goapp serve hellofoo/app


To run hellobar on dev server

    goapp serve hellofoo/app


if you want to run both apps on dev servers with different ports

    goapp serve -port=9090 -admin_port=9091 hellobar/app
    
### Trace test

Let's see the trace information on google cloud console.
Call the following url via 'GET' method, so that the call from hellofoo fetchs the hellobar's url.
Then, the information should be shown on the trace console.

    http://localhost:8080/trace-from-foo-to-bar