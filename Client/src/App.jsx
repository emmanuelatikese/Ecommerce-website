import { Route, Routes } from 'react-router-dom'
import SignUp from "./pages/signUp"

function App() {

  return (

      <div className='w-screen h-screen'>
        <Routes>
        <Route element={<SignUp/>} path='/SignUp'/>
        </Routes>
      </div>


  )
}

export default App
