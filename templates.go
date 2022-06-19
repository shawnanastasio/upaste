/**
 * Copyright 2019 Shawn Anastasio
 *
 * This file is part of upaste.
 *
 * upaste is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * upaste is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with upaste.  If not, see <https://www.gnu.org/licenses/>.
 */

/**
 * This file contains HTML templates.
 */

package main

// Index page served to curl/wget (plain text)
const IndexTemplateText =
`upaste(1)                          UPASTE                          upaste(1)

NAME
    upaste: simple command line pastebin.

SYNOPSIS
    <command> | curl -F 'upaste=<-' {{.ServerURL}}

DESCRIPTION
    upaste is a dead-simple pastebin with zero dependencies,
    meant to be used from the terminal.

    To upload from a browser, go here:
    {{.ServerURL}}/upload

    This server is configured to delete pastes after {{.ExpiryDays}} days.

EXAMPLES
    $ echo "Hello, upaste!" | curl -F 'upaste=<-' {{.ServerURL}}
        Upload the text "Hello, upaste!" to this server

    $ curl {{.ServerURL}} | less
        Display this page in a terminal

AUTHOR
    Written by Shawn Anastasio.
    Full source available at: https://github.com/shawnanastasio/upaste
`

// Index page served to browsers (HTML)
const IndexTemplateHTML =
`<html>
<head>
<title>upaste</title>
<body>
<pre>
upaste(1)                          UPASTE                          upaste(1)

NAME
    upaste: simple command line pastebin.

SYNOPSIS
    &lt;command&gt; | curl -F 'upaste=<-' {{.ServerURL}}

DESCRIPTION
    upaste is a dead-simple pastebin with zero dependencies,
    meant to be used from the terminal.

    <a href="/upload">To upload from a browser, go here.</a>

    This server is configured to delete pastes after {{.ExpiryDays}} days.

EXAMPLES
    $ echo "Hello, upaste!" | curl -F 'upaste=<-' {{.ServerURL}}
        Upload the text "Hello, upaste!" to this server

    $ curl {{.ServerURL}} | less
        Display this page in a terminal

AUTHOR
    Written by Shawn Anastasio.
    <a href="https://github.com/shawnanastasio/upaste">Full source available here.</a>
</pre>
</body>
`

// Upload page served to browsers (HTML)
const UploadTemplateHTML =
`<html>
<head>
<meta charset="UTF-8">
<title>Submit paste</title>
</head>
<body>
<form action="/" method="POST" accept-charset="UTF-8">
    <textarea name="upaste" cols="110" rows="30"></textarea>
    <br>
    <button type="submit">Upload</button>
</form>
</body>
</html>
`
