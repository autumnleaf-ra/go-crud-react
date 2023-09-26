import { useState } from 'react'
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import List from './users/List'
import ViewUser from './users/ViewUser'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
     <Router>
        <Routes>
          <Route index element={<List/>}/>
          <Route path='/profile/:id' element={<ViewUser/>}/>
        </Routes>
     </Router>
    </>
  )
}

export default App
