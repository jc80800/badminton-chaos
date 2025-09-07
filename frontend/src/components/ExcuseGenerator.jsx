import { useState } from 'react'

function ExcuseGenerator() {
  const [excuse, setExcuse] = useState('')
  const [loading, setLoading] = useState(false)

  const generateExcuse = async () => {
    setLoading(true)
    try {
      const response = await fetch('http://localhost:8080/random')
      const data = await response.json()
      setExcuse(data.message)
    } catch (error) {
      setExcuse('Failed to generate excuse. Try again!')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div>
      <button 
        onClick={generateExcuse}
        disabled={loading}
        className="button"
      >
        {loading ? 'Generating...' : 'Generate Excuse'}
      </button>

      {excuse && (
        <div className="excuse-display">
          "{excuse}"
        </div>
      )}
    </div>
  )
}

export default ExcuseGenerator
