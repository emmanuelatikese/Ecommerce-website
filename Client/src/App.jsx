import { Route, Routes } from "react-router-dom";
import NavBar from "./components/nav"
import Home from './pages/home';

function App() {

  return (
    <div className="w-full p-5 h-screen bg-color1 text-writeColor">
    <NavBar/>
    <Routes>
    <Route path="/Home" element={Home}/>
    </Routes>
    </div>
  )
}

export default App
