import { Route, Routes } from 'react-router-dom'
import SignUp from "./pages/signUp"
import Login from "./pages/login"
import Home from "./pages/home"
import Nav from "./components/nav"
import AdminDashboard from './pages/adminDashboard'

function App() {

  return (

      <div className='w-full h-auto p-0 m-0 bg-dusk-100 flex flex-col'>
      <Nav/>
        <Routes>
        <Route element={<SignUp/>} path='/SignUp'/>
        <Route element={<Login/>} path='/Login'/>
        <Route element={<Home/>} path='/'/>
        <Route element={<AdminDashboard/>} path='/dashboard' />
        </Routes>
      </div>


  )
}

export default App
