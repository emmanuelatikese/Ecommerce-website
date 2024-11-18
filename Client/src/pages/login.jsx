import { FaUser } from "react-icons/fa";
import { RiLockPasswordFill } from "react-icons/ri";
import { Link } from "react-router-dom";
import { IoLogInSharp } from "react-icons/io5";



const Login = () => {
  return (
    <div className="flex justify-center">
    
    <div className="flex flex-col items-center p-10 w-96  bg-white shadow-md rounded-md absolute top-36 gap-10">
    <h1 className="font-custom text-color3 text-2xl">Login</h1>

    <div className="flex flex-col gap-10">

        <div className="w-72 bg-color4 flex flex-row gap-4 rounded-md shadow-inner p-4 items-center">
        <FaUser className="text-color3"/>
            <input type="text" placeholder="Username" className="outline-none bg-color4" required />
        </div>

        <div className="w-72 bg-color4 flex flex-row gap-4 rounded-md shadow-inner p-4 items-center">
        <RiLockPasswordFill className="text-color3"/>
            <input type="password" placeholder="Password" className="outline-none bg-color4" required/>
        </div>


    </div>
     <Link className="">
        <button className="bg-color2 text-white p-2 w-36 shadow-md rounded font-bold font-custom flex flex-row items-center gap-2"> <IoLogInSharp className="text-white w-8 h-8"/> <p>Login</p></button>
     </Link>
    </div>
    
    </div>
  )
}

export default Login