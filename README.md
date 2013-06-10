go-ebay
=======

Experiments with the eBay API using golang

Currently, search-listings.go is extremely simplistic. All it does is return the first 
10 (or less if there aren't 10 results) eBay listings from the provided search queries. 
This program searches 5 different eBay sites at the moment.

**Updates**  
6/2/13 - Item price, shipping price, and seller location are now displayed  
6/3/13 - The new default is to search 5 eBay sites with 5 queries. Default site array contains eBay US, France, Germany, Italy, and Spain. Default query array contains corresponding translations of the Magic: The Gathering card Goblin Grenade. These searches operate concurrently.  
6/4/13 - Minor comment additions  
6/5/13 - Added items to To Do list  
6/8/13 - Displays locations that seller is shipping to  
6/9/13 - Displays listing URL, Item image thumbnail accessible in XMLResponse  

**To Use**  
You must make the following alterations in the code:  
1. Assign your eBay production AppID to appID  
2. Put your search terms into query array   
3. (optional) Specify how many sites to query with for loop in main()  
4. (optional) Adjust amount of search results [per site] with num_items
  

**To Do**
- flesh out Item struct to allow access to all XML response information
- allow users to easily specify which eBay site is being queried[1]
- accomodate for advanced query options such as filtering and sorting
- general error handling/fool-proofing, etc.  
- ~~Provide support for foreign language characters (some listing titles display improperly)~~[2]  
- Provide for [optional] automatic translation for cross-site queries (instead of specifying translations manually in query array)  
- Print more information regarding auctions (start and end times, number of bids)  
- Make it easier to filter between Auction/Buy it Now/Auction with BIN
 
[1] eBay sites can be specified with site array. 5 default sites are specified. Items in 
 query array should correspond to sites in site array. I'll work on a more user-friendly implementation.  
[2] Go 1.1 seems to have fixed many character issues. However, the black star (U+2605) character displays as "?", and I currently
 have not found a fix for that.