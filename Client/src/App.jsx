import {Route, Router} from "react-router-dom"
import SignUp from "./pages/signUp"
import Login from "./pages/login"
import Home from "./page/home"

function App() {

  return (
    <>
     
    <Router>
    <Route element={SignUp} path="/signUp"/>
    <Route element={Login} path="/login" />
    <Route element={Home} path="/home"/>
    </Router>
    </>
  )
}

export default App
