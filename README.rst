############################
Set Inline Breakpoints in Go
############################

.. Design-by-contract assert mechanism and breakpoint funnel.
.. Golang Inline breakpoint mechanism and Assertion Tool
.. Developer ``assert`` and breakpoint setter for Go

Problem(s):

* You miss the ability to set breakpoints *in* code
  (like Python's ``pdb.set_trace()``
  or perhaps C's ``std::raise(SIGINT);``).

* (You might also miss Design-by-Contract, a/k/a using soft asserts to catch developer errors).

Solution:

* Use
  `Derek Parker
  <http://derkthedaring.com/>`__'s
  awesome
  `Delve debugger for Go
  <https://github.com/derekparker/delve>`__.

* Use this project's simple soft assert mechanism to catch
  contract failures and get dumped at a Delve debug prompt.

==============
Docker Example
==============

Suppose you have an application that exposes a service on port 3000,
and that you would like to debug it.

An (abbreviated) Dockerfile might look like this:

.. code-block:: Dockerfile

    FROM golang:1.6

    ...

    EXPOSE 3000 3001
    ENV GOPATH /app
    ENTRYPOINT ["/app/bin/dlv", "debug", "--headless=true", "--listen=0.0.0.0:3001", "--log", "my-app"]

Build the container:

.. code-block:: bash

    GOPATH=$GOPATH:$(pwd -P) go build my-app

Run the container (and use ``--privileged`` so that ``dlv`` works):

.. code-block:: bash

    docker run -d --privileged -p 3000:3000 -p 3001:3001 --name my-app my-app

Connect to the debugger:

.. code-block:: bash

    dlv connect localhost:3001

Hook the ``Contract()`` function::

    (dlv) break contract.go:19

And then don't forget to run your app with ``continue``::

    (dlv) c

To set breakpoints in your application code, add soft asserts, like this:

.. code-block:: go

    import "github.com/landonb/golang-contract"

    contract.Contract(false)

When this code fires, your ``dlv`` terminal will plop you at a prompt.

=====
Hints
=====

It's easy to set the breakpoint
-------------------------------

Fortunately, you don't have to type or copy-paste the ``break contract.go:19``
command each time. Much like you can up-arrow in Bash to see previous
commands, so can also do so you in Delve.

- And to make it even more convenient, you can narrow the search with
  a leading initial character.

  E.g., type ``b`` and the hit the up arrow, then probably Enter -- c'est tout.::

    (dlv) b<UP><ENTER>

It's easy to ``tail`` the container output
------------------------------------------

You can easily tail the container log without using ``docker logs``
and having to restart the log every time you rebuild the container.

- Send the container's stdout to ``rsyslogd`` and ``tail -f`` a real file.

  For example, start the ``syslog`` daemon, and then point docker to it.

Create a configuration file. We'll call it ``rsyslogd.conf``:

.. code-block:: bash

    cat > /path/to/my/syslog.conf << EOF
        $ModLoad imtcp
        $InputTCPServerRun 10514
        *.* /path/to/my/syslog.log
    EOF

- Run the daemon. Point it to your file and also to a PID file that it'll maintain:

.. code-block:: bash

    touch /path/to/my/syslog.pid
    /usr/sbin/rsyslogd \
        -f /path/to/my/syslog.conf \
        -i /path/to/my/syslog.pid

- And then from another terminal, tail it:

.. code-block:: bash

    tail -f /path/to/my/syslog.log

- Now run your container, and specify the ``syslog`` logging driver:

.. code-block:: bash

    docker run -d --privileged -p 3000:3000 -p 3001:3001 \
        --log-driver syslog \
        --log-opt syslog-address=tcp://localhost:10514 \
        --name my-app my-app

It's easy to restart the container
----------------------------------

Every time you exit from the debugger, your application halts. It's easy to restart it:

.. code-block:: bash

    docker restart my-app

You'll probably need 3 terminal windows to work effectively
-----------------------------------------------------------

- One terminal connected to the Delve server.

- A second terminal tailing the container log.

- And a third terminal you'll use, e.g., to build
  and run the container, and to test it (say, by
  using ``curl`` to tickle the API on port 3000).

Note that you won't have to touch the ``tail`` terminal,
but anytime you rebuild and rerun the ``my-app`` container,
you'll have to ``Ctrl-D`` from the debugger,
``dlv connect`` again, and set the break again
(using the trick described above).

