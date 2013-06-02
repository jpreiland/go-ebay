go-ebay
=======

Experiments with the eBay API using golang

Currently, search-listings.go is extremely simplistic. All it does is return the first 
10 eBay listings from the provided search query.

To Use, you must make the following alterations in the code:
1. Assign your eBay production AppID to appID
2. Put your search terms into query
3. (optional) Adjust amount of search results with num_items

to do:
- flesh out Item struct to allow access to all XML response information
- provide support for multiple concurrent searches
- allow users to easily specify which eBay site is being queried
- accomodate for advanced query options such as filtering and sorting
- general error handling/fool-proofing, etc.