go-ebay
=======

Experiments with the eBay API using golang

Currently, search-listings.go is extremely simplistic. All it does is return the first 
10 eBay listings from the provided search query.

**Updates**  
6/2/13 - Item price, shipping price, and seller location are now displayed  
6/3/13 - The new default is to search 5 eBay sites with 5 queries. Default site array contains eBay US, France, Germany, Italy, and Spain. Default query array contains corresponding translations of the Magic: The Gathering card Goblin Grenade. These searches should operate concurrently (assuming I coded it correctly)  
6/4/13 - Minor comment additions

**To Use**  
You must make the following alterations in the code:  
1. Assign your eBay production AppID to appID  
2. Put your search terms into query array   
3. (optional) Specify how many sites to query with for loop in main() (limit to one iteration to only search US eBay site)  
4. (optional) Adjust amount of search results [per site] with num_items
  

**To Do**
- flesh out Item struct to allow access to all XML response information
- provide support for multiple concurrent searches[1]
- allow users to easily specify which eBay site is being queried[2]
- accomodate for advanced query options such as filtering and sorting
- general error handling/fool-proofing, etc.  
- Provide support for foreign language characters (some listing titles display improperly)  
- Provide for [optional] automatic translation for cross-site queries (instead of specifying translations manually in query array)

[1] multiple searches implemented, not sure exactly how effective the concurrency is  
[2] eBay sites can be specified with site array. 5 default sites are specified. Items in 
 query array should correspond to sites in site array. I'll work on a more user-friendly implementation.