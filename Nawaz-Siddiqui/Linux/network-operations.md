# Network Operations

## wget

**wget** is a command line utility that can capably handle:

Large file downloads

Recursive downloads, where a web page refers to other web pages and all are downloaded at once

Password-required downloads

Multiple file downloads.

## curl

Besides downloading, you may want to obtain information about a URL, such as the source code being used. curl can be used from the command line or a script to read such information. curl also allows you to save the contents of a web page to a file, as does wget.

### Commands

| Command | Function |
| --- | --- |
| wget <url> | To download a web page |
| curl <URL> | To get the content of a webpage |
| curl -o saved.html http://www.mysite.com | To get the contents of a web page and store it to a file named saved.html |
