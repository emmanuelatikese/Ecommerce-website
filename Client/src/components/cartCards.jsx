import Trouser from "../assets/images/trouser.jpg"
import { MdDelete } from "react-icons/md";
import { IoIosAddCircle } from "react-icons/io";
import { BiSolidMinusCircle } from "react-icons/bi";

const CartCards = () => {
  return (

      
      <div className=" p-4 round-md shadow-md flex flex-row gap-12">
      <div className="bg-color4 p-4 shadow-md rounded-md"><img src={Trouser} className="w-36 h-36"/></div>
      <div className="bg-color4 p-4 flex flex-row gap-8 rounded-md shadow-md">
        <div className="flex flex-col gap-2 p-2 justify-center ">
        <p className="font-custom ">Trousers</p>
        <p className="font-bold text-gray-600">Cloth</p>
        <MdDelete className="w-6 h-6 text-red-600"/>
        </div>

        <div className="flex flex-row gap-10 justify-center items-center">
          <div  className="flex flex-row justify-center items-center gap-4">
            <IoIosAddCircle className="w-8 h-8 shadow-md text-color3 cursor-pointer"/>
            <p className="text-lg font-bold">1</p>
            <BiSolidMinusCircle className="w-8 h-8 shadow-md text-color3 cursor-pointer"/>
          </div>

          <p className="font-bold text-lg text-green-700">$180</p>
        </div>
      </div>
      </div>
      
  )
}

export default CartCards
