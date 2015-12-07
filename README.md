Go smtpd [![GoDoc](https://godoc.org/github.com/strothj/smtpd?status.png)](https://godoc.org/github.com/strothj/smtpd)
========

Package smtpd implements an SMTP server in golang.

Features
--------

* STARTTLS (using `crypto/tls`)
* Authentication (PLAIN/LOGIN, only after STARTTLS)
* XCLIENT (for running behind a proxy)
* Connection, HELO, sender and recipient checks for rejecting e-mails using callbacks
* Configurable limits for: connection count, message size and recipient count
* Hands incoming e-mail off to a configured callback function

Original Repositories
---------------------
The original source repositories can be found here:
[https://bitbucket.org/chrj/smtpd](https://bitbucket.org/chrj/smtpd)
[https://bitbucket.org/kardianos/smtpd](https://bitbucket.org/kardianos/smtpd)

Feedback
--------

I'm merging in changes from the above mentioned repositories, the original author
can be reached at [email](mailto:christian@technobabble.dk).
