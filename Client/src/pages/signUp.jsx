import { FaUser } from "react-icons/fa6";

const signUp = () => {
  return (
    <div className="flex flex-col gap-8 w-80 h-96 justify-center items-center mx-auto my-auto"> 
    <h3>SignUp</h3>
    
    <div >
    <div className="flex flex-row gap-3">
    <FaUser/>
    <input type="text" placeholder="username" className="w-5"/>
    </div>

    <div className="flex flex-row gap-3">
    <FaUser/>
    <input type="text" placeholder="username" className="w-5"/>
    </div>

    </div>

    </div>
  )
}

export default signUp