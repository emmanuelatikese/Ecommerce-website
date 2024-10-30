import { Route, Routes } from 'react-router-dom'
import SignUp from "./pages/signUp"

function App() {

  return (

      <div className='flex h-7 w-5'>
        <Routes>
        <Route element={<SignUp/>} path='/SignUp'/>
        </Routes>
      </div>


  )
}

export default App
