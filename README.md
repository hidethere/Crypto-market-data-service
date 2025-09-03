# 📈 Crypto Market Data Service  

## 🌍 Overview  

`Crypto-market-data-service` is a distributed **real-time crypto ticker service** built in Go.  

The service evolved from a simple **REST API** into a **WebSocket-based streaming system** that can:  
- 🔄 Fetch live ticker data from **Binance**.  
- 🎯 Allow clients to subscribe to **all symbols** or **specific cryptos** (e.g., `BTCUSDT`, `ETHUSDT`).  
- ⚡ Handle **dynamic subscriptions** (clients can update filters on the fly).  
- 🛡️ Maintain robust connections with **ping/pong heartbeat** and **automatic reconnection**.  

Clients can interact with the service using either:  
- **REST API** endpoints (`/api/v1/ticker?symbol=...`) for on-demand queries.  
- **WebSockets** (`/ws/tickers`) for real-time streaming.  

---

## ⚙️ Functionalities  

### 🔹 REST API – Get Latest Ticker(s)  
**Request:**  
```http
GET /api/v1/ticker?symbol=BTCUSDT
GET /api/v1/ticker?symbol=BTCUSDT,ETHUSDT
````
**Response (single symbol):**
```http
{
  "symbol": "BTCUSDT",
  "bid": "40450.25",
  "ask": "40451.10",
  "last": "40450.70",
  "high": "41000.00",
  "low": "39800.00",
  "volume": "1234.56"
}
````
**Response (multiple symbols):**
````json
{
   [
    {
      "symbol": "BTCUSDT",
      "bid": "40450.25",
      "ask": "40451.10",
      "last": "40450.70",
      "high": "41000.00",
      "low": "39800.00",
      "volume": "1234.56"
    },
    {
      "symbol": "ETHUSDT",
      "bid": "3045.10",
      "ask": "3045.80",
      "last": "3045.50",
      "high": "3200.00",
      "low": "2950.00",
      "volume": "9876.54"
    }
  ]
}
````

### 🔹 WebSocket – Real-Time Streaming
Default Mode:
- When a client connects to (`/ws/tickers`), it receives all tickers from Binance in real-time.

Dynamic Subscription Update:
- Clients can filter which symbols they want to receive using messages like:
  ```json
  { "subscribe": ["BTCUSDT", "ETHUSDT"] }
  
- To reset back and receive **all tickers again**, send:
  ```json
  { "subscribe": [] }
  ````

## 🚀 How to Run

### 🔹Run Locally
Install dependencies:
````
go mod tidy
````
Run the service:
```
go run main.go
````
### 🔹Run with Docker Compose
If you prefer using docker-compose, simply run:
```
docker-compose up --build
````
This will:
- Build the service image.
- Start the service on port 8080
