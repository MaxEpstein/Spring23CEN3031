# Sprint 2

## Completed Issues
    Updated About Page
    Created functioning search bar & page
    Includeed logo and link to home page in navbar
    Upon login, redirects to dashboard
    Started configuring dashboard
    Integrated frontend and backend
    User can search and see latest price for stock through search bar



## Frontend tests
    Nav bar test using cypress framework:
        Tests navigation from home to search
        Tests navigation from search to home
        Tests Login Functionality
        Tests Login redirect to dashboard

# Backend tests

    main.go 

   Server for the search page on the front end.Using a
   websocket protocol, the server is waiting for 
   the front end search button to be pressed in order to
   search for the necessary ticker, set up all list, and then 
   return the current price.

    main_workfunc_list.go
    
   This is the main file for collecting all the necessary data
   
    add_historic_data(temp_stock *stock)
This function adds the historic data (timestamp of open and close,with corresponding price)
from the beggining of the stocks history

    setup_main_working_list(s_type_name []string, s_type_sym []string) *data_list 
This will collect a list of strings of stock types(etfs, stocks, etc) along with 
corresponding tickers and create the main container that will 
hold all stocks. An empty container can also be initialized with 
nil values. Returns data_list.

The stock container:
>map[string]map[string]stock



    getDataByTicker(ticker string, s_type string) *stock 

This function is mainly to populate all the different attributes 
of each stock. Returns a stock pointer. 

    update_data_list(working_list *data_list)

This function adds the current value to each stock in the list along with 
current time stamp. Returns nothing.

    addStockToMain(stockToAdd *stock, main_list *data_list)

Adding the individual stock to the main working list. Returns nothing.

    checkIfStockExist(ticker string) bool 

Checks to see if the ticker actually exist within the stock market
to keep the program from accessing invalid pointer.
