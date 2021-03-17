The purpose of the simpleGrep example application is to illustrate how the Go regexp package can be used to search files for text strings or patterns using regular expressions.

To run this application, you might need to enter "go get github.com/ncruces/zenity" (in addition to "go get github.com/zserge/lorca") in a terminal.

The github.com/ncruces/zenity repository offers a Go package for using dialogs. The simpleGrep application uses it for displaying a folder chooser dialog or file chooser dialog when the user clicks on the "Select Folder" or "Select File" buttons respectively. SimpleGrep also uses it for displaying warning and error messages (e.g. when the user clicks on "Search" without entering both a pattern and a path.

You can enter plain text or a regular expression into the pattern input box. For example to find the word "Box", enter "Box". To find any word that starts with a 'B' enter "\b[B]". Only a few simple regular expressions have been tested, so it is probably best not to depend on the results when entering complex regular expressions. For case insensitive, whole word, or whole line searches, you can check the corresponding checkboxes.  So instead of entering "(?i)box" to find either "Box" or "box", you can enter either one and click on the "Case Insensitive" checkbox.

You can use the "Select Folder" or "Select File" buttons to populate the path input box. Alternatively, you type the path directly. You can enter a file name without entering the full path if it is in the current directory. You can also enter "." to search for all files in the current directory.

If path is a directory, then simpleGrep will search that directory recursively.

For a grep-like application for production use consider sift at "https://github.com/svent/sift". It is a command line application (as opposed to a GUI application), but is fast and powerful.  It is written in Go.

