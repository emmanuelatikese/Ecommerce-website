import { Route, Routes } from 'react-router-dom'
import SignUp from "./pages/signUp"

function App() {

  return (

      <div className='w-full h-screen border-black'>
      <p>App container</p>
        <Routes>
        <Route element={<SignUp/>} path='/SignUp'/>
        </Routes>
      </div>


  )
}

export default App
