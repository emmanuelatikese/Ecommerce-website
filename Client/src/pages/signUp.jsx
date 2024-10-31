import { FaEnvelope, FaLock, FaUser } from "react-icons/fa6";
import { Link } from "react-router-dom";
import { IoIosPersonAdd } from "react-icons/io";

const signUp = () => {
  return (
    <div className="p-8 w-96 h-auto absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 text-center border-solid border-4 border-marigold-200 flex flex-col gap-8 "> 
    <h1 className="text-marigold-200 text-2xl font-bold">SignUp</h1>
    
    <div className="flex flex-col gap-6">
      <div className="flex flex-row gap-3 p-3  flex-wrap justify-center items-center">
      <FaUser className=" text-marigold-200 w-8 h-8"/>
      <input type="text" required placeholder="Username" className="border-b-marigold-200 border-b-4 p-1 w-60 outline-none bg-dusk-100  font-san font-bold text-marigold-200"/>
      </div>

      <div className="flex flex-row gap-3 p-3  flex-wrap justify-center items-center">
      <FaEnvelope className=" text-marigold-200 w-8 h-8"/>
      <input type="text" required placeholder="Email" className="border-b-marigold-200 border-b-4 p-1 w-60 outline-none bg-dusk-100  font-san font-bold text-marigold-200"/>
      </div>

      <div className="flex flex-row gap-3 p-3  flex-wrap justify-center items-center">
      <FaLock className=" text-marigold-200 w-8 h-8"/>
      <input type="password" required placeholder="Password" className="border-b-marigold-200 border-b-4 p-1 w-60 outline-none bg-dusk-100  font-san font-bold text-marigold-200"/>
      </div>

      <div className="flex flex-row gap-3 p-3  flex-wrap justify-center items-center">
      <FaLock className=" text-marigold-200 w-8 h-8"/>
      <input type="password" required placeholder="Confirm Password" className="border-b-marigold-200 border-b-4 p-1 w-60 outline-none bg-dusk-100  font-san font-bold text-marigold-200"/>
      </div>
      <p className="text-marigold-200 font-bold text-sm"> You already have an Account? Click <Link to={"/Login"} className="underline text-white">Here</Link></p>
      
      <div className="flex flex-row justify-center w-80 pb-4  pt-4  flex-wrap">
      <button className=" rounded-2xl text-white bg-marigold-200 font-bold p-2 w-32">
        <Link className="flex flex-row flex-wrap justify-center items-center gap-2"> <IoIosPersonAdd className="w-5 h-5"/>SignUp</Link>
      </button>

      </div>

</div>

    </div>
  )
}

export default signUp