# Plutus
Tracking prices and volume for stocks and determining critical points, serving as buy/sell signals.


## TODOs
- [ ] Add DB to store the data & set up repo and schema
- [ ] Use Finnhub API to get data (Price, Volume & moving average)
- [ ] Store the data from Finnhub in the DB
- [ ] Define TODOs for the calculations



## v1 Requirements
- Track price of when bought & when sold
- Track price action (call for data every 5 mins)
    - Identify trends, by comparing with last hours
    - Identify price ceilings and floors (resistance) & when they are broken
    - If during breaking of the resistance the trading volume is > 2x the average volume → extra buy signal??

- Provide moving average for a stock of
    - 10 days
    - 50 days
    - 100 days

- Track Volume spikes
    - Volume > 3x average is significant
        - Average over (14/50/100 days)

### Determining the line of least resistance

- Price movement and breaking of pivotal resistances (or big numbers like first 100 break)
    
    → Determine resistances and alert if broken
    
    → Alert on big number breaks
    
- Volume matters, strong volume indicates strong conviction
    
    → Monitor volume and alert if it’s higher than normal, especially combined with price action
    
- Tape Reading
    - Rapid, consistent price ticks upward with increasing trades means buyers are dominant - up is the path.
    - Hesitation or erratic moves signal resistance
- Trend context
    - In bull market, upward breakouts are the path of least resistance
    - In bear market, breakdowns are the path of least resistance
