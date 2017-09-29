package main

/*
Simulator specifies all the mechanisms that aid in environment randomization and by extension agent generalization
to the real world such as:
	- Periodically changing format of responses.
	- Periodically changing enabled assets and pairs
	- Periodically changing allowed order types
	- Periodically changing ticker interval
	- Periodically changing the utilization measure of the agent smith.
	- Periodically changing config such as fees
	- Periodically changing the matching engine used
	- Periodically changing enabled functionality
	- Periodically performing random errors
	- The starting amount and margin amount
	- Margin and funding parameters
	- Systematically creating agents
	- .etc

	Analytics combine simulated events with account performance .etc

Endpoints:
     Public:
	- Ticker
	- Pairs
	- Pairs detailed
	- Assets
	- Trades
	- Orderbook
	- Fundingbook
	- Lends
	- 24 hour stats
	- Historic Rates
	- Time

     Private:
     	- Accounts
     	- Account Fees
     	- Balances
     	- Fills
     	- Position
     	- Deposits

     	- Withdrawals
     	- Withdrawal History

     	- Transfer
     	- Transfer History

     	- New Order
     	- Cancel Order
     	- Cancel All Orders
     	- Cancel Session Orders
     	- Active Orders
     	- Order History
     	- Order Status

     	- New Offer
     	- Cancel Offer
     	- Offers
     	- Active Credits
     	- Funding trades
     	- Active Funding used
     	- Active funding not used

     	- Positions

How to emulate with different endpoints and changing endpoints?

*/
