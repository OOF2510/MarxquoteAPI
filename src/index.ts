import { Hono } from 'hono'
import quotes from '../quotes.json'

const app = new Hono()

app.get('/', (c) => {
  const randomIndex = Math.floor(Math.random() * quotes.quotes.length)
  return c.text(quotes.quotes[randomIndex].text + ' - ' + quotes.quotes[randomIndex].author)
})

app.get('/getquotejson', async (c) => {
  const randomIndex = Math.floor(Math.random() * quotes.quotes.length)
  return c.json(quotes.quotes[randomIndex])
})

export default app
