The purpose of the database example application is to illustrate how the Go "database/sql" and "github.com/go-sql-driver/mysql" packages can be used to perform SQL queries on either a mysql or mariadb database.

The community edition of Mysql can be downloaded from https://dev.mysql.com/downloads/. Mariadb can be downloaded from https://mariadb.org/download/.

This application uses the nation sample database. To load this database you can use either the nation.sql file, which is in this database directory (the "database" directory), or download "nation.zip" from "https://www.mariadbtutorial.com/getting-started/mariadb-sample-database/" and extract the nation.sql file from the nation.zip file. From a terminal change the directory to where the nation.sql file is located. Then connect to the mysql or mariadb server by entering "mysql -u root -p". After entering the root password enter "source c:\mariadb\nation.sql" in order to load the nation database.

The database application only uses the "countries" table. Furthermore, it only uses the "name", "area", and "national_day" columns from that table.

To use the application, enter the database password into the "Database password" input box and then click on the "Connect to database" button. If you enter a incorrect password, the application will not report the error at this time.  However, when you subsequently click on the "Query" button, the application will display a error message dialog and then close when you dismiss the dialog.  Regardless whether the password is correct or incorrect, after you click on the "Connect to database" button it should become disabled and the "Query" button should become enabled.

To use the "Exact match" option, ensure that the respective radio button is selected, enter the full name of a country into the "Name of country" input box and click on the "Query" button.  For example, if you entered "Canada", the following should appear in the text box:

"Name                                              Area           National Day
-----------------------------------------------------------------------------
Canada                                            9970610.00     1867-07-01
-----------------------------------------------------------------------------

To use the "Match on starting letter(s)" option, ensure that the respective option is selected, enter one or more starting letters of a country's name, and click on the "Query" button.  For example, if you entered "Ta", the following should appear in the text box:

Name                                              Area           National Day
-----------------------------------------------------------------------------
Tajikistan                                        143100.00      1991-09-09
-----------------------------------------------------------------------------
Taiwan                                            36188.00       1911-10-10
-----------------------------------------------------------------------------
Tanzania                                          883749.00      1961-12-09
-----------------------------------------------------------------------------

