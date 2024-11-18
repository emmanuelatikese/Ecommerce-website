import { FaUser } from "react-icons/fa";
import { RiLockPasswordFill } from "react-icons/ri";
import { MdEmail } from "react-icons/md";
import { Link } from "react-router-dom";



const SignUp = () => {
  return (
    <div className="flex justify-center">
    
    <div className="flex flex-col items-center p-10 w-96  bg-white shadow-md rounded-md absolute top-36 gap-10">
    <h1 className="font-custom text-color3 text-2xl">SignUp</h1>

    <div className="flex flex-col gap-10">

        <div className="w-72 bg-color4 flex flex-row gap-4 rounded-md shadow-inner p-4 items-center">
        <FaUser className="text-color3"/>
            <input type="text" placeholder="Username" className="outline-none bg-color4" required />
        </div>

        <div className="w-72  bg-color4 flex flex-row gap-4 rounded-md shadow-inner p-4 items-center">
        <MdEmail className="text-color3"/>
            <input type="text" placeholder="Email" className="outline-none bg-color4" required/>
        </div>

        <div className="w-72 bg-color4 flex flex-row gap-4 rounded-md shadow-inner p-4 items-center">
        <RiLockPasswordFill className="text-color3"/>
            <input type="password" placeholder="Password" className="outline-none bg-color4" required/>
        </div>

        <div className="w-72 bg-color4 flex flex-row gap-4 rounded-md shadow-inner p-4 items-center">
        <RiLockPasswordFill className="text-color3"/>
            <input type="password" placeholder="Confirm Password" className="outline-none bg-color4" required/>
        </div>

    </div>
     <Link>
        <button className="bg-color2 text-white p-4 w-36 shadow-md rounded"> SignUp</button>
     </Link>
    </div>
    
    </div>
  )
}

export default SignUp