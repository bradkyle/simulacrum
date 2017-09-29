The engine is where all the market functionality for the simulacrum occurs

Engine:
    - Has Ticker Endpoints
    - Has Orderbook Endpoints
    - Has Trade Endpoints
    - Has Stats Endpoints
    - Has Margin Funding
        - Broker
        - Fundingbook


Order Matching Engines:
    - Continuous Trading:
        - FifoTwoTradeEngine
        - FifoTwoTradeLMMEngine
        - FifoTwoTradeTopLMMEngine
        - FifoOneTradeEngine
        - FifoOneTradeLMMEngine
        - FifoOneTradeTopLMMEngine
        - FifoOneTradeTopLMMEngine
        - ProRataEngine
        - ThresholdProRataEngine
        - ThresholdProRataLMMEngine
        - AllocationEngine
    - Auction Trading:
        -

Order Execution Types:
    - MarginMarket
	- MarginLimit
	- MarginStop
	- MarginTrailingStop
	- MarginFillOrKill
	- ExchangeMarket
	- ExchangeLimit
	- ExchangeStop
	- ExchangeTrailingStop
	- ExchangeFillOrKill
	- BrokerMarket
	- BrokerLimit
	- BrokerStop
	- BrokerTrailingStop
	- BrokerFillOrKill

Assets:
    - Allowed Assets

Pairs (Market)
    - Open/Closed
    - Allowed Asset Pairs
    - Fees
    - Has Ticker Endpoint
    - Has Stats Endpoint
    - Has Orderbook Endpoint
    - Min/Max order

Ticker:
    - Ticker frequency
    - Ticker fields




