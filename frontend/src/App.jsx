import React, { useEffect, useState } from 'react'
import axios from 'axios'
import './App.css'

const API = import.meta.env.VITE_API_URL || 'http://localhost/api'

export default function App() {
  const [plans, setPlans] = useState([])

  useEffect(() => {
    axios.get(`${API}/plans`)
      .then(res => setPlans(res.data))
      .catch(() => {})
  }, [])

  return (
    <div className="container">
      <header className="hero">
        <h1>🗾 Trip Road</h1>
        <p>みんなの観光プランを共有しよう</p>
      </header>

      <main>
        <h2>新着プラン</h2>
        {plans.length === 0 ? (
          <p className="empty">まだプランがありません</p>
        ) : (
          <div className="plan-list">
            {plans.map(p => (
              <div key={p.id} className="plan-card">
                <h3>{p.title}</h3>
                <p>{p.description}</p>
                <span className="badge">{p.prefecture}</span>
                <span className="badge">{p.days}日間</span>
              </div>
            ))}
          </div>
        )}
      </main>
    </div>
  )
}
