import { Component, createResource } from 'solid-js'

const App: Component = () => {
  const [data, { refetch }] = createResource(
    () => ['data'],
    async () => {
      const res = await fetch('http://localhost:4000')
      return await res.json()
    }
  )

  setInterval(() => refetch(), 1000)

  return (
    <>
      <main
        style={{
          width: '100vw',
          height: '100vh',
          display: 'flex',
          'flex-direction': 'column',
          'align-items': 'center',
          'justify-content': 'center',
          'background-color': 'black',
          color: 'white'
        }}
      >
        <h1>Moonrepo</h1>

        <pre>
          <code>{JSON.stringify(data(), null, 2)}</code>
        </pre>
      </main>
    </>
  )
}

export default App
