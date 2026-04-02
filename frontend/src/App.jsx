import React, { useEffect, useState } from 'react'
import axios from 'axios'

const API = import.meta.env.VITE_API_URL || 'http://localhost/api'

export default function App() {
  const [plans, setPlans] = useState([])
  const [status, setStatus] = useState('loading')

  useEffect(() => {
    axios.get(`${API}/health`)
      .then(() => setStatus('ok'))
      .catch(() => setStatus('error'))

    axios.get(`${API}/plans`)
      .then(res => setPlans(res.data))
      .catch(() => {})
  }, [])

  return (
    <div style={{ maxWidth: 800, margin: '0 auto', padding: 24 }}>
      <h1>Trip Road</h1>
      <p>観光プラン投稿サイト</p>
      <p>API Status: <strong>{status}</strong></p>
      <h2>プラン一覧</h2>
      {plans.length === 0 ? (
        <p>まだプランがありません</p>
      ) : (
        <ul>
          {plans.map(p => (
            <li key={p.id}>{p.title} - {p.prefecture} ({p.days}日間)</li>
          ))}
        </ul>
      )}
    </div>
  )
}
