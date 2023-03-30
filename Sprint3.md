# Sprint 3

## Completed Issues
    - The Graph on the Search page is now customizable to specific time periods for the searched stock
    - Can get all data within a period of time given that period of time and a time interval
    - Skip or pass weekends when finding a starting date of a time interval
    - Graph buttons and switching between time intervals 
    - Added Saved Stocks section to redirect back to search page
    - Add save to dashboard function frontend
    - Bug testing for search and dashboard

## Incomplete Issues
    - Getting saved stocks to redirect to desired graph
    - Saving user information to a database

## Backend Functions
    main.go
        Server for the search page on the front end. Using a
        websocket protocol, the server is waiting for 
        the front end search button to be pressed in order to
        search for the necessary ticker, set up all list, and then 
        return the current price.

    worklistFunctions.go
        This is the main file for collecting all the necessary data. 
        Returns data to frontend based on strings passed by the frontend.

        Functions:
        - initializeWorkingList(s_type_name []string, s_type_sym []string, data_interval string, data_time_interval string) *data_list
            This will collect a list of strings of stock types(etfs, stocks, etc) along with 
            corresponding tickers and create the main container that will 
            hold all stocks. An empty container can also be initialized with 
            nil values. Returns data_list.
        - addHistoricalData(temp_stock *stock, timeFrame string, chartInterval string) 
            This function adds the historic data to a map with a key value equal to a UNIX 
            uint64, and a map value of the price * 100. This map is passed to the frontend
            to make the graphs.
        - getTimeFrame(timeFrame string, chartIntervalString string) (*datetime.Datetime, datetime.Interval)
            Takes in two strings from the frontend to determine the time period and the chart
            interval to be used in the graph.
        - passWeekends(numDays int) time.Time
            For use in 1day and 5day intervals, makes sure weekends are passed over
        - skipWeekends(numDays int) time.Time
            For use in all other time periods other than 1day and 5day,
            makes sure weekends are skipped
        - getChartInterval(chartIntervalString string) datetime.Interval
            Uses case statement to determine interval for use in graph creation
            based on passed in string from frontend.
        - getDataByTicker(ticker string, s_type string, data_interval string, data_time_interval string) *stock
            This function is mainly to populate all the different attributes 
            of each stock. Returns a stock pointer. 
        - updateMainWorkingList(working_list *data_list)
            This function adds the current value to each stock in the list along with 
            current time stamp. Returns nothing.
        - addStockToMain(stockToAdd *stock, main_list *data_list)
            Adding the individual stock to the main working list. 
            Returns nothing.
        - checkIfStockExist(ticker string) bool 
            Checks to see if the ticker actually exists within the stock market
            to keep the program from accessing an invalid pointer.
        -func HashPassword(password string) (string, error)
            Generates a hashed string for password encryption
        -func CheckPasswordHash(password, hash string) bool 
            Returns the hash code back into a regular password. 
        -func createEncryptedInfo(username string, pw string) string 
            Helper function to create the inital encrypted message to be stored.
        -func userNew(username string, pw string) *user 
            Constructor for the user and their stocks. This contains a username, password, list of 
            tickers, and an encrypted string that saves once the program is complete. 
        -func addTicker(currentUser *user, ticker string) 
            To add a ticker to the users list, and add it to the encryption

## Frontend Tests
    Nav bar test using cypress framework:
        Tests navigation from home to search
        Tests navigation from search to home
    Tests Login Functionality
    Tests Login redirect to dashboard
    Graph Change Tests

## Backend Tests
    Sprint 1/2:
        testInitializeWorkingList
        testGetDataByTicker
        testAddStockToMain
        testCheckIfStockExist
    Sprint 3:
        testGetTimeFrame
        testPassWeekends
        testSkipWeekends
        testAddHistoricalData



