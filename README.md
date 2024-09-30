# Cryptocurrency Trading Microservices

This project consists of multiple microservices designed to facilitate cryptocurrency trading data analysis and management, built primarily with Golang, PostgreSQL, and Redis.

## Services

### 1. Crawler

The Crawler service is responsible for fetching candle data from Binance and storing it in a PostgreSQL database.

**Key Features:**
- Connects to Binance API
- Retrieves candle data for specified cryptocurrencies
- Stores data in PostgreSQL database

**Tech Stack:**
- Golang
- PostgreSQL
- Binance API

### 2. Trader View

The Trader View service provides an interface to query candle data from the database.

**Key Features:**
- Queries candle data from PostgreSQL
- Exposes API endpoints for data retrieval

**Tech Stack:**
- Golang
- PostgreSQL

### 3. Backend

The Backend service handles user authentication, management, mock payments, and performs various technical analysis calculations. It also implements response caching for improved performance.

**Key Features:**
- User authentication and management
- Mock payment system
- Retrieves candle data from Trader View service via NATS.io
- Calculates technical indicators:
  - Moving Average (MA)
  - Simple Moving Average (SMA)
  - Exponential Moving Average (EMA)
  - Double Exponential Moving Average (DEMA)
  - Relative Strength Index (RSI)
  - Bollinger Bands
  - Moving Average Convergence Divergence (MACD)
  - Stochastic Oscillator
- Implements response caching in Redis for candle data and trading view algorithm results

**Tech Stack:**
- Golang
- NATS.io for inter-service communication
- Redis for caching

## Tech Stack Overview

- **Programming Language:** Golang
- **Database:** PostgreSQL
- **Caching:** Redis
- **Message Broker:** NATS.io


