# Passphrase Generator app

This repository is the fist app I make using [Go language](http://www.golang.org/).


### Start the web server:

    revel run myapp

   Run with <tt>--help</tt> for options.

### Go to http://localhost:9000/ and you'll see:

"It works"

### The passphrase generator

The objective of this app is to generate a safe passphrase from the user selected common phrase, such as **"MyNameIsJohnAndIWasBornOn1968"**.

The algorithm can be reviewed at this other [repository](https://github.com/LeonardoCastro/MyFirstGoCodes/tree/master/password). **No personal information nor generated passphrases are beheld.**

### Description of Contents

The default directory structure of a generated Revel application:

    myapp               App root
      app               App sources
        controllers     App controllers
          init.go       Interceptor registration
        models          App domain models
        routes          Reverse routes (generated code)
        views           Templates
      tests             Test suites
      conf              Configuration files
        app.conf        Main configuration file
        routes          Routes definition
      messages          Message files
      public            Public assets
        css             CSS files
        js              Javascript files
        images          Image files

app

    The app directory contains the source code and templates for your application.

conf

    The conf directory contains the application’s configuration files. There are two main configuration files:

    * app.conf, the main configuration file for the application, which contains standard configuration parameters
    * routes, the routes definition file.


messages

    The messages directory contains all localized message files.

public

    Resources stored in the public directory are static assets that are served directly by the Web server. Typically it is split into three standard sub-directories for images, CSS stylesheets and JavaScript files.

    The names of these directories may be anything; the developer need only update the routes.

test

    Tests are kept in the tests directory. Revel provides a testing framework that makes it easy to write and run functional tests against your application.

## Revel

This appe was made using [Revel](revel.github.io).
