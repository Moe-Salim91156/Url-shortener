## URL shortener Requirements :

	## 1. accept HTTP requests , (API)
	## 2. takes a Long URL , and genereate a short one linked to it 
	## 3. when clicked on the generated URL it should REDIRECT the user to the actual url



## UNITS:

	## v2 : 
		## added Users (models + store(interface))
		## auth service
		## Users owning Urls , Create, list, Delete URls
		## auth middleware?
		## UI , GO templates


## General Architectural Flow concluded : 
	## request 
	## middleware
	## -> handler (translation)
	## -> service (desicions)
	## -> store (contracts/interfaces , HOW to Data)
	## -> Data Layer (WHAT DATA)
