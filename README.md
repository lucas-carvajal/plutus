# Plutus
Tracking prices and volume for stocks and determining critical points, serving as buy/sell signals.


## TODOs
- [X] Use TwelveData API to get data (Price, Volume & moving average)
- [ ] Store the data from TwelveData in the DB
- [ ] Define TODOs for the calculations
- [ ] Define further data to calc EOD and store in other tables
- [ ] Integrate with Telegram
- [ ] Conditions for notification sending
- [ ] Add tests



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

# Technical Set Up

### Run local DB with docker
Create postgres container with volume
```
docker run --name postgres-db \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=yourpassword \
  -e POSTGRES_DB=quotes_db \
  -p 5432:5432 \
  -v pgdata:/var/lib/postgresql/data \
  -d postgres:latest
```

Connect to the DB
```
psql -h localhost -U postgres -d quotes_db
```

Connection string for local DB
```
postgres://postgres:yourpassword@localhost:5432/quotes_db?sslmode=disable
```

### Run the server
```
go run main.go
```