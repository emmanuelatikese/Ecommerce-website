import { Route, Routes } from "react-router-dom";
import NavBar from "./components/nav"
import Home from "./pages/home"
import SignUp from "./pages/signUp";

function App() {

  return (
    <div className="w-full p-5 h-full bg-color1 text-writeColor flex flex-col">
    <NavBar/>
    <Routes>
    <Route path="/Home" element={<Home/>}/>
    <Route path="/SignUp" element={<SignUp/>} />
    </Routes>
    </div>
  )
}

export default App
