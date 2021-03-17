The purpose of the mimicDatabase example application is to illustrate the look and feel of the database application without using a database. For information as to how the database application behaves, please see the README.md file in the database folder.

The mimicDatabase application uses the data from the countries.txt file, which is located in this directory (the mimicDatabase directory). This file was created by executing a "The SELECT ... INTO OUTFILE" statement on the nation database. Whereas the database application get its data from the database, the mimicDatabase application get its data from the countries.txt file. Like the database application, the mimicDatabase application only uses the "name", "area", and "national_day" data from the countries table.

In the mimicDatabase application, the password input box is automatically filled with the text "password". If you change it and then click on the "Connect to database" button, the application will close and log a error message to the terminal.

The mimicDatabase application performs case-sensitive searches; whereas, the database application uses case-insensitive searches.
