import { Route, Routes } from 'react-router-dom'
import SignUp from "./pages/signUp"

function App() {

  return (

      <div className='w-full h-screen p-0 m-0 bg-dusk-100'>
        <Routes>
        <Route element={<SignUp/>} path='/SignUp'/>
        </Routes>
      </div>


  )
}

export default App
