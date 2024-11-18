import { Route, Routes } from "react-router-dom";
import NavBar from "./components/nav"
import Home from "./pages/home"
import SignUp from "./pages/signUp";
import Login from "./pages/login";

function App() {

  return (
    <div className="w-full p-5 h-full bg-color1 text-writeColor flex flex-col">
    <NavBar/>
    <Routes>
    <Route path="/Home" element={<Home/>}/>
    <Route path="/SignUp" element={<SignUp/>} />
    <Route path="/Login" element={<Login/>} />
    </Routes>
    </div>
  )
}

export default App
